package api

type Handlers struct {
	UserHandler        *UserHandler
	ProductHandler     *ProductHandler
	TransactionHandler *TransactionHandler
	AuthHandler        *AuthHandler
	ShipmentHandler    *ShipmentHandler
	OrderHandler       *OrderHandler
	PackageHandler     *PackageHandler
	OrderLineHandler   *OrderLineHandler
	TierListHandler    *TierListHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler, TransactionHandler *TransactionHandler, AuthHandler *AuthHandler, ShipmentHandler *ShipmentHandler, OrderHandler *OrderHandler, PackageHandler *PackageHandler, OrderLineHandler *OrderLineHandler, TierListHandler *TierListHandler) *Handlers {
	return &Handlers{
		UserHandler:        UserHandler,
		ProductHandler:     ProductHandler,
		TransactionHandler: TransactionHandler,
		AuthHandler:        AuthHandler,
		ShipmentHandler:    ShipmentHandler,
		OrderHandler:       OrderHandler,
		PackageHandler:     PackageHandler,
		OrderLineHandler:   OrderLineHandler,
		TierListHandler:    TierListHandler,
	UserHandler              *UserHandler
	ProductHandler           *ProductHandler
	TransactionHandler       *TransactionHandler
	AuthHandler              *AuthHandler
	ShipmentHandler          *ShipmentHandler
	OrderHandler             *OrderHandler
	PackageHandler           *PackageHandler
	OrderLineHandler         *OrderLineHandler
	SupplierHandler          *SupplierHandler
	SupplierOrderListHandler *SupplierOrderListHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler,
	TransactionHandler *TransactionHandler, AuthHandler *AuthHandler,
	ShipmentHandler *ShipmentHandler, OrderHandler *OrderHandler,
	PackageHandler *PackageHandler, OrderLineHandler *OrderLineHandler,
	SupplierHandler *SupplierHandler, SupplierOrderListHandler *SupplierOrderListHandler) *Handlers {
	return &Handlers{
		UserHandler:              UserHandler,
		ProductHandler:           ProductHandler,
		TransactionHandler:       TransactionHandler,
		AuthHandler:              AuthHandler,
		ShipmentHandler:          ShipmentHandler,
		OrderHandler:             OrderHandler,
		PackageHandler:           PackageHandler,
		OrderLineHandler:         OrderLineHandler,
		SupplierHandler:          SupplierHandler,
		SupplierOrderListHandler: SupplierOrderListHandler,
	}
}
