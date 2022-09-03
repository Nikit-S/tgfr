package template

type Screen struct {
	Elems       []IBlock
	SkipOnStart bool
	Repeat      bool
	OnElement   int
}

//main function for execute screen
//screen execeutes all its elements one by one
//if element exec returns false it decides to proceed or finish
//execution stops on endMsg | gotoScreen | nil element
func (s Screen) Execute(bot *Bot, user *User) {
	//defer user.SetScreen(nil)
	//user.SetScreen(s)
	if s.SkipOnStart {
		<-user.GetChan()
	}
	if s.Elems == nil {
		return
	}
	s.OnElement = user.onElement
	for s.OnElement < len(s.Elems) {
		s.execNext(bot, user)
	}
}

func (s *Screen) execNext(bot *Bot, user *User) {
	if len(s.Elems) == 0 {
		return
	}
	blocks := s.Elems[user.onElement]
	user.onElement++
	blocks.Execute(bot, user)
	s.OnElement++
}

func (s Screen) GetFirstElement() IBlock {
	if len(s.Elems) == 0 {
		return nil
	}
	return s.Elems[0]
}

//func (s Screen) NextScreen *Screen {
//	return s.NextScreen
//}
