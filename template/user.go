package template

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	userId     int64
	active     bool
	updChan    chan tgbotapi.Update
	lastUpdate tgbotapi.Update
	onScreen   *Screen
	onElement  int
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

func (u *User) GetLastUpdate() tgbotapi.Update {
	return u.lastUpdate
}

func (u *User) SetLastUpdate(update tgbotapi.Update) {
	u.lastUpdate = update
}

func (u *User) OnScreen() *Screen {
	return u.onScreen
}

func (u *User) OnElement() int {
	return u.onElement
}

func (u *User) GetUserId() int64 {
	return u.userId
}

func (u *User) SetScreen(screen *Screen) {
	u.onScreen = screen
}

func (u *User) SetElement(element int) {
	u.onElement = element
}
