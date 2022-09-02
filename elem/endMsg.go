package elem

import (
	"github.com/Nikit-S/tgfr/template"
)

type EndMsg struct {
}

func (m EndMsg) Execute(bot *template.Bot, user *template.User) {
	user.SetScreen(nil)
	user.SetElement(-1)

}
