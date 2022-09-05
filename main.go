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
			Text: "Hello, User! You' ii wait now for 3 seconds. Can try to spam",
		},
		blocks.Timer{
			Seconds: 5,
		},
		blocks.Msg{
			Text: "I m bot! For next 4 messages i will repeat your message. Or use command to stop me from main line.",
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
		blocks.Msg{
			Text: "You ll wait now for 3 seconds. Can try to spam",
		},
		blocks.Timer{
			Seconds: 3,
		},

		blocks.Msg{
			Text: "STOP I will repeat your message",
		},
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
