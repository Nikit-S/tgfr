package blocks

import (
	"fmt"

	"github.com/Nikit-S/tgfr/template"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Msg struct {
	Text string
}

func (m Msg) Execute(bot *template.Bot, user *template.User) (exit bool) {
	msg := tgbotapi.NewMessage(user.GetUserId(), m.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}
	return false
}

type RepeatInput struct {
	Text string
	Num  int
}

func (m RepeatInput) Execute(bot *template.Bot, user *template.User) (exit bool) {
	update := <-user.GetChan()
	msg := tgbotapi.NewMessage(user.GetUserId(), update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}
	return false
}

func (m RepeatInput) GetNum() int {
	return m.Num
}
