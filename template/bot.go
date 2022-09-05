package template

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Main stucrcture of bot
type Bot struct {
	*tgbotapi.BotAPI
	users     map[int64]*User
	storage   IStorage
	workspace *Workspace
}

//returns new bot with initialized storage
func NewBot(key string, w *Workspace) *Bot {
	bt, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		panic(err)
	}

	return &Bot{BotAPI: bt, users: make(map[int64]*User), workspace: w}
}

//starts the mainLoop of updates receiving
func (b Bot) Start() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.GetUpdatesChan(updateConfig)
	b.mainLoop(updates)
	return nil
}

//do nothing
func (b *Bot) Stop() error {
	return nil
}

func (b *Bot) UserLoop(userId int64) {
	user := b.GetUser(userId)
	for {
		select {
		case block := <-user.blockChan:
			block.Execute(b, user)
			user.onElement++
			go b.GetWorkspace().SendNextElementToUser(user)
		case command := <-user.commandChan:
			command.Execute(b, user)
			user.onCommand++
			go b.GetWorkspace().SendNextCommandElementToUser(user)
		}
	}
}

func (b *Bot) mainLoop(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if b.GetUser(update.Message.From.ID) == nil {
			user := NewUser(update.Message.From.ID)
			b.users[update.Message.From.ID] = user
			go b.UserLoop(user.userId)
			screen, ok := b.workspace.GetScreen(user.onScreen)
			if !ok {
				//user.Reset()
			}
			user.blockChan <- screen.GetElement(user.onElement)

		}
		b.GetUser(update.Message.From.ID).SetLastUpdate(update)
		b.routeUpdate(update)
	}
}

func (b *Bot) routeUpdate(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	user := b.GetUser(update.Message.From.ID)
	if update.Message.IsCommand() {
		user.onCommand = 0
		user.onElement = -2
		commStr := update.Message.Command()
		var comm Command
		var ok bool
		if _, ok = user.activeCommands[commStr]; !ok {
			return false
		}
		if comm, ok = b.workspace.GetCommand(commStr); !ok {
			return false
		}
		select {
		case <-user.waitChan:
		default:
			user.CloseWaitChan()
		}
		user.commandChan <- comm.GetFirstElement()
	} else {
		select {
		case <-user.waitChan:
		case user.updChan <- update:
		}
	}
	return true
}

func (b *Bot) GetAllUsers() (map[int64]*User, error) {
	if b.users == nil {
		return nil, fmt.Errorf("users not initialized")
	}
	return b.users, nil
}

func (b *Bot) GetUser(UserId int64) *User {
	if _, ok := b.users[UserId]; !ok {
		return nil
	} else {
		return b.users[UserId]
	}
}

//func (b *Bot) CreateUser(UserId int64, screen *Screen) *User {
//	user := &User{
//		userId:   UserId,
//		active:   true,
//		onScreen: screen,
//		updChan:  make(chan tgbotapi.Update),
//		state:    STATE_WAIT_INPUT,
//	}
//	if screen.GetFirstElement() != nil {
//		user.onElement = 0
//	} else {
//		panic("first screen has no elements")
//	}
//	b.users[UserId] = user
//	return user
//}

//
//func (b *Bot) StartDialog(UserId int64, screen *Screen) {
//	user := b.GetUser(UserId)
//	screen.Execute(b, user)
//}

func (b *Bot) GetApi() *tgbotapi.BotAPI {
	return b.BotAPI
}

func (b *Bot) GetStorage() IStorage {
	return b.storage
}

func (b *Bot) SetStorage(storage IStorage) {
	b.storage = storage
}

func (b *Bot) GetWorkspace() *Workspace {
	return b.workspace
}
