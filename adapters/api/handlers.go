package api

type Handlers struct {
	UserHandler        *UserHandler
	ProductHandler     *ProductHandler
	TransactionHandler *TransactionHandler
	ShipmentHandler    *ShipmentHandler
	OrderHandler       *OrderHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler, TransactionHandler *TransactionHandler, ShipmentHandler *ShipmentHandler, OrderHandler *OrderHandler) *Handlers {
	return &Handlers{
		UserHandler:        UserHandler,
		ProductHandler:     ProductHandler,
		TransactionHandler: TransactionHandler,
		ShipmentHandler:    ShipmentHandler,
		OrderHandler:       OrderHandler,
	}
}
