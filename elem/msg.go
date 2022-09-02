package elem

import (
	"fmt"

	"github.com/Nikit-S/tgfr/template"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Msg struct {
	Text string
}

func (m Msg) Execute(bot *template.Bot, user *template.User) {
	msg := tgbotapi.NewMessage(user.GetUserId(), m.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}

}

type RepeatMsg struct {
	Text string
	Num  int
}

func (m RepeatMsg) Execute(bot *template.Bot, user *template.User) {
	fmt.Println("repeat msg wait", user.GetChan())
	update := <-user.GetChan()
	fmt.Println("repeat msg wait passed", user.GetChan())
	msg := tgbotapi.NewMessage(user.GetUserId(), update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}

}

func (m RepeatMsg) GetNum() int {
	return m.Num
}
