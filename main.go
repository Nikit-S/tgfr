package main

import (
	"log"
	"os"

	"github.com/Nikit-S/tgfr/elem"
	"github.com/Nikit-S/tgfr/template"
)

func main() {

	bot := template.NewBot(
		os.Getenv("TG_TOKEN"),
		RepeatScreen,
	)
	log.Fatal(bot.Start())
}

var HelloScreen = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IElement{
		elem.Msg{
			Text: "Hello, User!",
		},
		elem.Msg{
			Text: "I m bot!",
		},
		elem.EndMsg{},
	},
}

var SecondScreen = template.Screen{
	Elems: []template.IElement{
		elem.Msg{
			Text: "What!",
		},
		elem.Msg{
			Text: "it works!",
		},
	},
}

var RepeatScreen = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IElement{
		elem.RepeatMsg{},

		elem.GotoScreen{
			Screen: &HelloScreen,
		},
	},
}
