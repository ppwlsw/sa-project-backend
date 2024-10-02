package api

type Handlers struct {
	UserHandler        *UserHandler
	ProductHandler     *ProductHandler
	TransactionHandler *TransactionHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler, TransactionHandler *TransactionHandler) *Handlers {
	return &Handlers{
		UserHandler:        UserHandler,
		ProductHandler:     ProductHandler,
		TransactionHandler: TransactionHandler,
	}
}
