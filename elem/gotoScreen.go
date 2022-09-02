package elem

import (
	"github.com/Nikit-S/tgfr/template"
)

type GotoScreen struct {
	Screen *template.Screen
}

func (m GotoScreen) Execute(bot *template.Bot, user *template.User) {
	user.SetScreen(m.Screen)
	user.SetElement(0)

}
