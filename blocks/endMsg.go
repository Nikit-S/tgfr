package blocks

import (
	"github.com/Nikit-S/tgfr/template"
)

type EndMsg struct {
}

// Set user screen to nil and and element to 0 which means that bot will stop redirecting updates to this user
func (m EndMsg) Execute(bot *template.Bot, user *template.User) {
	user.SetElement(-2)
}

func (m EndMsg) String() string {
	return "EndMsg"
}
