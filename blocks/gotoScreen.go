package blocks

import (
	"github.com/Nikit-S/tgfr/template"
)

type GotoScreen struct {
	Screen  *template.Screen
	Element int
}

func (m GotoScreen) Execute(bot *template.Bot, user *template.User) (exit bool) {
	user.SetScreen(m.Screen)
	user.SetElement(m.Element)

	go user.OnScreen().Execute(bot, user)
	return true
}
