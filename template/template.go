package template

//type IBot interface {
//	Start() error
//	Stop() error
//	GetAllUsers() (map[int64]IUser, error)
//	GetUser(UserId int64) IUser
//	CreateUser(UserId int64, screen IScreen) IUser
//	StartDialog(UserId int64, screen IScreen)
//	GetApi() *tgbotapi.BotAPI
//}
//
//type IUser interface {
//	Activate()
//	IsActive() bool
//	Deactivate()
//	GetChan() chan tgbotapi.Update
//	GetUserId() int64
//	GetLastUpdate() tgbotapi.Update
//	OnScreen() IScreen
//	SetScreen(screen IScreen)
//	OnElement() IBlock
//	SetElement(element IBlock)
//}
//
//type WorkSpace struct{}
//
//type IScreen interface {
//	Execute(bot IBot, user IUser)
//	GetFirstElement() IBlock
//	NextScreen() IScreen
//}

type IBlock interface {
	Execute(bot *Bot, user *User)
	String() string
}

type INotification interface {
	Send() error
}
