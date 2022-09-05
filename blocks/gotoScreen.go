package blocks

import (
	"fmt"

	"github.com/Nikit-S/tgfr/template"
)

type GotoScreen struct {
	Screen  string
	Element int
}

// Set user screen to m.Screen and and element to m.Element and then starts executing new screen in new goroutine
func (m GotoScreen) Execute(bot *template.Bot, user *template.User) {
	user.SetScreen(m.Screen)
	user.SetElement(m.Element - 1)

	screen, ok := bot.GetWorkspace().GetScreen(m.Screen)
	if !ok {
		return
	}
	elem := screen.GetElement(m.Element)
	if elem == nil {
		return
	}
}

func (m GotoScreen) String() string {
	return fmt.Sprintf("GotoScreen: %s[%d]", m.Screen, m.Element)
}
