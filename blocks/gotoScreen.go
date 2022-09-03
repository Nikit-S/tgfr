package blocks

import (
	"github.com/Nikit-S/tgfr/template"
)

type GotoScreen struct {
	Screen  *template.Screen
	Element int
}

// Set user screen to m.Screen and and element to m.Element and then starts executing new screen in new goroutine
func (m GotoScreen) Execute(bot *template.Bot, user *template.User) {
	user.SetScreen(m.Screen)
	user.SetElement(m.Element)

	go user.OnScreen().Execute(bot, user)
}
