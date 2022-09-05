package blocks

import (
	"time"

	"github.com/Nikit-S/tgfr/template"
)

//todo таймер сна

//скорее всего нужно заменить временно в юзере переменные экрана и элемента на пустые, пустить горутину со слипом и по ее завершению вернуть переменные обратно не забыть про канал который можкт ждать

type Timer struct {
	Seconds int
}

func (m Timer) Execute(bot *template.Bot, user *template.User) {
	screen := user.OnScreen()
	element := user.OnElement()
	user.SetScreen("")
	user.SetElement(-2)
	openChan := false
	select {
	case <-user.GetWaitChan():
		openChan = true
	default:
		user.CloseWaitChan()
	}
	go func() {
		time.Sleep(time.Duration(m.Seconds) * time.Second)
		user.SetScreen(screen)
		user.SetElement(element + 1)
		if openChan {
			user.OpenWaitChan()
		}
		go bot.GetWorkspace().SendNextElementToUser(user)
	}()

}

func (m Timer) String() string {
	return "Timer"
}
