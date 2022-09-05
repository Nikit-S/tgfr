package template

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	userId         int64
	active         bool
	updChan        chan tgbotapi.Update
	lastUpdate     tgbotapi.Update
	onScreen       string
	onCommand      int
	onElement      int
	activeCommands map[string]struct{}
	blockChan      chan IBlock
	commandChan    chan IBlock
	waitChan       chan struct{}
}

func NewUser(UserId int64) *User {
	user := &User{
		userId:      UserId,
		active:      true,
		onScreen:    STARTSCREEN,
		onElement:   0,
		updChan:     make(chan tgbotapi.Update),
		waitChan:    make(chan struct{}),
		blockChan:   make(chan IBlock),
		commandChan: make(chan IBlock),
	}
	user.CloseWaitChan()
	return user
}

func (u *User) Activate() {
	u.active = true
}

func (u *User) IsActive() bool {
	return u.active
}

func (u *User) Deactivate() {
	u.active = false
}

func (u *User) GetChan() chan tgbotapi.Update {
	return u.updChan
}

func (u *User) GetWaitChan() chan struct{} {
	return u.waitChan
}

func (u *User) GetLastUpdate() tgbotapi.Update {
	return u.lastUpdate
}

func (u *User) SetLastUpdate(update tgbotapi.Update) {
	u.lastUpdate = update
}

func (u *User) OnScreen() string {
	return u.onScreen
}

func (u *User) OnElement() int {
	return u.onElement
}

func (u *User) GetUserId() int64 {
	return u.userId
}

func (u *User) SetScreen(screen string) {
	u.onScreen = screen
}

func (u *User) SetElement(element int) {
	u.onElement = element
}

func (u *User) SetCommands(Commands map[string]struct{}) {
	u.activeCommands = Commands
}

func (u *User) GetCommands() map[string]struct{} {
	return u.activeCommands
}

func (u *User) OpenWaitChan() {
	u.waitChan = make(chan struct{})
}

func (u *User) CloseWaitChan() {
	close(u.waitChan)
}

func (u *User) RecieveUpdate() (tgbotapi.Update, bool) {
	var update tgbotapi.Update
	select {
	case <-u.GetWaitChan():
		return update, false
	case update = <-u.GetChan():
		return update, true
	}
}

func (u *User) SendUpdate(update tgbotapi.Update) {
	u.updChan <- update
}

func (u *User) SendExec(exec IBlock) {
	u.blockChan <- exec
}

func (u *User) RecieveCommand() (IBlock, bool) {
	var command IBlock
	select {
	case <-u.GetWaitChan():
		return command, false
	case command = <-u.commandChan:
		return command, true
	}
}

func (u *User) SendCommand(command IBlock) {
	u.commandChan <- command
}

func (u *User) RecieveExec() (IBlock, bool) {
	var exec IBlock
	select {
	case exec = <-u.blockChan:
		return exec, true
	case <-u.GetWaitChan():
		return exec, false
	}
}

const (
	STATE_NONE       = "none"
	STATE_WAIT_INPUT = "wait_input"
)
