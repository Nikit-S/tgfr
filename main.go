package main

import (
	"log"
	"os"

	"github.com/Nikit-S/tgfr/blocks"
	"github.com/Nikit-S/tgfr/template"
)

func main() {
	w := template.NewWorkspace()
	w.AddScreen(HelloScreen)
	w.AddScreen(RepeatScreen)
	w.AddCommand(RepeatCommand)
	bot := template.NewBot(
		os.Getenv("TG_TOKEN"),
		w,
	)
	log.Fatal(bot.Start())
}

// Screen for greetings and further ending
var HelloScreen = template.Screen{
	Name:        template.STARTSCREEN,
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.Msg{
			Text: "Hello, User!",
		},
		blocks.Msg{
			Text: "I m bot!",
		},
		blocks.GotoScreen{
			Screen:  "RepeatScreen",
			Element: 0,
		},
		//blocks.EndMsg{},
	},
}

//
//var SecondScreen = template.Screen{
//	Elems: []template.IBlock{
//		blocks.Msg{
//			Text: "What!",
//		},
//		blocks.Msg{
//			Text: "it works!",
//		},
//	},
//}
//
//repeats last user input and redirects to RepeatScreen2
var RepeatCommand = template.Command{
	Name:        "repeat",
	Description: "Repeats last user input",
	Chain: []template.IBlock{
		blocks.RepeatInput{},
		blocks.EndMsg{},
	},
}

var RepeatScreen = template.Screen{
	Name:        "RepeatScreen",
	SkipOnStart: false,
	Repeat:      false,
	Elems: []template.IBlock{
		blocks.SetCommand{
			Commands: map[string]struct{}{
				"repeat": {},
			},
		},
		blocks.RepeatInput{},
		blocks.RepeatInput{},
		blocks.RepeatInput{},
		blocks.RepeatInput{},
		blocks.Msg{
			Text: "By, User!",
		},
		blocks.EndMsg{},
	},
}

//var RepeatScreen2 = template.Screen{
//	SkipOnStart: false,
//	Repeat:      false,
//	Elems: []template.IBlock{
//		blocks.RepeatInput{},
//
//		blocks.GotoScreen{
//			Screen: &RepeatScreen3,
//		},
//		//blocks.EndMsg{},
//	},
//}
//
//var RepeatScreen3 = template.Screen{
//	SkipOnStart: false,
//	Repeat:      false,
//	Elems: []template.IBlock{
//		blocks.RepeatInput{},
//
//		blocks.GotoScreen{
//			Screen: &HelloScreen,
//		},
//		//blocks.EndMsg{},
//	},
//}
