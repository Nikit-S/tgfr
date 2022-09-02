package template

import "fmt"

type Screen struct {
	Elems       []IElement
	SkipOnStart bool
	Repeat      bool
}

func (s *Screen) Execute(bot *Bot, user *User) {
	//user.SetScreen(s)
	fmt.Println("exec screen")
	if s.SkipOnStart {
		<-user.GetChan()
	}
	for user.onElement < len(s.Elems) {

		//fmt.Println("user", user)
		if !s.execNext(bot, user) {
			break
		}
		//time.Sleep(1 * time.Second)
	}
	fmt.Println("screen end", user.OnScreen())
	if s.Repeat {
		fmt.Println("repeat")
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
