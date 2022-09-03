package template

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Main stucrcture of bot
type Bot struct {
	*tgbotapi.BotAPI
	users       map[int64]*User
	startScreen Screen
	storage     IStorage
}

//returns new bot with initialized storage
func NewBot(key string, StartScreen Screen) *Bot {
	bt, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		panic(err)
	}

	return &Bot{BotAPI: bt, users: make(map[int64]*User), startScreen: StartScreen}
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

func (b *Bot) mainLoop(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if b.GetUser(update.Message.From.ID) == nil {
			user := b.CreateUser(update.Message.From.ID, &b.startScreen)
			go user.OnScreen().Execute(b, user)
		}
		b.GetUser(update.Message.From.ID).SetLastUpdate(update)
		go b.routeUpdate(update)
	}
}

func (b *Bot) routeUpdate(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	user := b.GetUser(update.Message.From.ID)
	if user != nil {
		if user.OnScreen() == nil {
			return false
		}
		b.GetUser(update.Message.From.ID).GetChan() <- update
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

func (b *Bot) CreateUser(UserId int64, screen *Screen) *User {
	user := &User{
		userId:   UserId,
		active:   true,
		onScreen: screen,
		updChan:  make(chan tgbotapi.Update),
	}
	if screen.GetFirstElement() != nil {
		user.onElement = 0
	} else {
		panic("first screen has no elements")
	}
	b.users[UserId] = user
	return user
}

//
//func (b *Bot) StartDialog(UserId int64, screen *Screen) {
//	user := b.GetUser(UserId)
//	screen.Execute(b, user)
//}

func (b *Bot) GetApi() *tgbotapi.BotAPI {
	return b.BotAPI
}
