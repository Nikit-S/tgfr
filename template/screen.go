package template

type Screen struct {
	Elems       []IElement
	SkipOnStart bool
	Repeat      bool
}

//main function for execute screen
//screen execeutes all its elements one by one
//if element exec returns false it decides to proceed or finish
//execution stops on endMsg | gotoScreen | nil element
func (s *Screen) Execute(bot *Bot, user *User) {
	//user.SetScreen(s)
	if s.SkipOnStart {
		<-user.GetChan()
	}
	for user.onElement < len(s.Elems) {
		if s.Elems[user.onElement] == nil {
			user.SetScreen(nil)
			user.SetElement(-1)
			return
		}

		if !s.execNext(bot, user) {
			break
		}
		//time.Sleep(1 * time.Second)
	}
	if s.Repeat {
		user.SetElement(0)
		user.SetScreen(s)
	}
	if user.OnScreen() != nil {
		defer user.OnScreen().Execute(bot, user)
	}

}

func (s *Screen) execNext(bot *Bot, user *User) bool {
	if len(s.Elems) == 0 {
		return false
	}
	elem := s.Elems[user.onElement]

	elem.Execute(bot, user)
	if user.onElement == -1 || user.onScreen != s {
		return false
	}
	user.onElement++
	if user.onElement >= len(s.Elems) {
		return false
	}
	return true
}

func (s Screen) GetFirstElement() IElement {
	if len(s.Elems) == 0 {
		return nil
	}
	return s.Elems[0]
}

//func (s Screen) NextScreen *Screen {
//	return s.NextScreen
//}
