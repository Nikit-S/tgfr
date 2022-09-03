package main

import (
	"log"
	"os"

	"github.com/Nikit-S/tgfr/blocks"
	"github.com/Nikit-S/tgfr/template"
)

func main() {

	bot := template.NewBot(
		os.Getenv("TG_TOKEN"),
		RepeatScreen,
	)
	log.Fatal(bot.Start())
}

// Screen for greetings and further ending
var HelloScreen = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.Msg{
			Text: "Hello, User!",
		},
		blocks.Msg{
			Text: "I m bot!",
		},
		blocks.EndMsg{},
	},
}

var SecondScreen = template.Screen{
	Elems: []template.IBlock{
		blocks.Msg{
			Text: "What!",
		},
		blocks.Msg{
			Text: "it works!",
		},
	},
}

//repeats last user input and redirects to RepeatScreen2
var RepeatScreen = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.RepeatInput{},

		blocks.GotoScreen{
			Screen: &RepeatScreen2,
		},
		//blocks.EndMsg{},
	},
}

var RepeatScreen2 = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.RepeatInput{},

		blocks.GotoScreen{
			Screen: &RepeatScreen3,
		},
		//blocks.EndMsg{},
	},
}

var RepeatScreen3 = template.Screen{
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.RepeatInput{},

		blocks.GotoScreen{
			Screen: &HelloScreen,
		},
		//blocks.EndMsg{},
	},
}
