package api

type Handlers struct {
	UserHandler        *UserHandler
	ProductHandler     *ProductHandler
	TransactionHandler *TransactionHandler
	AuthHandler        *AuthHandler
	ShipmentHandler    *ShipmentHandler
	OrderHandler       *OrderHandler
	PackageHandler     *PackageHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler, TransactionHandler *TransactionHandler, AuthHandler *AuthHandler, ShipmentHandler *ShipmentHandler, OrderHandler *OrderHandler, PackageHandler *PackageHandler) *Handlers {
	return &Handlers{
		UserHandler:        UserHandler,
		ProductHandler:     ProductHandler,
		TransactionHandler: TransactionHandler,
		AuthHandler:        AuthHandler,
		ShipmentHandler:    ShipmentHandler,
		OrderHandler:       OrderHandler,
		PackageHandler:     PackageHandler,
	}
}
