package blocks

import (
	"github.com/Nikit-S/tgfr/template"
)

type EndMsg struct {
}

func (m EndMsg) Execute(bot *template.Bot, user *template.User) (exit bool) {
	user.SetScreen(nil)
	user.SetElement(0)
	return false
}
