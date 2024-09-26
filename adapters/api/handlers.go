package api 

type Handlers struct {
	UserHandler *UserHandler
}

func ProvideHandlers(UserHandler *UserHandler) *Handlers {
	return &Handlers{
		UserHandler: UserHandler,
	}
}