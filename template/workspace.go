package template

type Workspace struct {
	Screens  map[string]Screen
	Commands map[string][]Command
}

func NewWorkspace() *Workspace {
	return &Workspace{
		Screens:  make(map[string]Screen),
		Commands: make(map[string][]Command),
	}
}

func (w *Workspace) AddScreen(screen Screen) {
	if _, ok := w.Screens[screen.Name]; ok {
		panic("screen with id " + screen.Name + " already exists")
	}
	w.Screens[screen.Name] = screen
}

func (w *Workspace) AddCommand(command Command) {
	w.Commands[command.Name] = append(w.Commands[command.Name], command)
}

func (w *Workspace) GetScreen(name string) (Screen, bool) {
	if _, ok := w.Screens[name]; !ok {
		return Screen{}, false
	}
	return w.Screens[name], true
}

func (w *Workspace) GetCommand(name string) (Command, bool) {
	if _, ok := w.Commands[name]; !ok {
		return Command{}, false
	}
	return w.Commands[name][0], true
}

func (w *Workspace) GetScreens() map[string]Screen {
	return w.Screens
}

func (w *Workspace) GetCommands() map[string][]Command {
	return w.Commands
}

func (w *Workspace) GetStartScreen() *Screen {
	if len(w.Screens) == 0 {
		panic("no screens in workspace")
	}
	if _, ok := w.Screens["startScreen"]; !ok {
		panic("no screen with id 0 in workspace")
	}
	screen := w.Screens["startScreen"]
	return &screen
}

func (w *Workspace) SendNextElementToUser(user *User) {
	screen, ok := w.GetScreen(user.OnScreen())
	if !ok {
		return
	}
	if user.OnElement() >= screen.GetLength() || user.OnElement() < 0 {
		return
	}
	elem := screen.GetElement(user.OnElement())
	if elem == nil {
		return
	}
	user.SendExec(elem)

}

func (w *Workspace) SendNextCommandElementToUser(user *User) {
	Command, ok := w.GetCommand(user.OnScreen())
	if !ok {
		return
	}
	if user.onCommand >= Command.GetLength() || user.onCommand < 0 {
		return
	}
	elem := Command.GetElement(user.onCommand)
	if elem == nil {
		return
	}
	user.SendCommand(elem)

}
