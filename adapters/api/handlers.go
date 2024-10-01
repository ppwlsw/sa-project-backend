package api

type Handlers struct {
	UserHandler    *UserHandler
	ProductHandler *ProductHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler) *Handlers {
	return &Handlers{
		UserHandler:    UserHandler,
		ProductHandler: ProductHandler,
	}
}
