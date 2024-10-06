package api

type Handlers struct {
	UserHandler        *UserHandler
	ProductHandler     *ProductHandler
	TransactionHandler *TransactionHandler
	AuthHandler        *AuthHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler, TransactionHandler *TransactionHandler, AuthHandler *AuthHandler) *Handlers {
	return &Handlers{
		UserHandler:        UserHandler,
		ProductHandler:     ProductHandler,
		TransactionHandler: TransactionHandler,
		AuthHandler:        AuthHandler,
	}
}
