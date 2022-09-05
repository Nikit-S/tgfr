package blocks

import (
	"github.com/Nikit-S/tgfr/template"
)

type SetCommand struct {
	Commands map[string]struct{}
}

// Sends message to user with text m.Text
func (m SetCommand) Execute(bot *template.Bot, user *template.User) {
	user.SetCommands(m.Commands)
}

func (m SetCommand) String() string {
	return "SetCommand"
}
