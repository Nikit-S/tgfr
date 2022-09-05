package template

//todo по аналогии с обработкой экрана нужно реализовать обработку команд
type Command struct {
	Name        string
	Description string
	Chain       []IBlock
}

func (c Command) Execute(bot *Bot, user *User) {
}

func (c Command) String() string {
	return "Command"
}

func (c Command) GetFirstElement() IBlock {
	if len(c.Chain) == 0 {
		return nil
	}
	return c.Chain[0]
}

func (c Command) GetLength() int {
	return len(c.Chain)
}

func (c Command) GetElement(index int) IBlock {
	if len(c.Chain) == 0 || index >= len(c.Chain) || index < 0 {
		return nil
	}
	return c.Chain[index]
}
