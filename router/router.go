package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/adapters/database"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"gorm.io/gorm"
)

func SetUpRouters(app *fiber.App, db *gorm.DB) {

	userRepo := database.ProvideUserPostgresRepository(db)
	userService := usecases.ProvideUserService(userRepo)
	userHandler := api.ProvideUserHandler(userService)

	productRepo := database.InitiateProductPostGresRepository(db)
	productService := usecases.InitiateProductsService(productRepo)
	productHandler := api.InitiateProductHandler(productService)

	transactionRepo := database.InitiateTransactionPostGresRepository(db)
	transactionService := usecases.InitiateTransactionService(transactionRepo)
	transactionHandler := api.InitiateTransactionHandler(transactionService)

	shipmentRepo := database.InitiateShipmentPostgresRepository(db)
	shipmentService := usecases.InitiateShipmentService(shipmentRepo)
	shipmentHandler := api.InitiateShipmentHandler(shipmentService)

	handlers := api.ProvideHandlers(userHandler, productHandler, transactionHandler, shipmentHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//User
	app.Post("/users", handlers.UserHandler.CreateUser)
	app.Get("/users", handlers.UserHandler.GetAllUsers)
	app.Get("/users/:id", handlers.UserHandler.GetUserByID)

	//Product
	app.Post("/products", handlers.ProductHandler.CreateProduct)
	app.Get("/products", handlers.ProductHandler.GetAllProducts)
	app.Get("product/:id", handlers.ProductHandler.GetProductByID)
	app.Put("products/:id", handlers.ProductHandler.UpdateProduct)

	//Transaction
	app.Post("/transaction", handlers.TransactionHandler.CreateTransaction)
	app.Post("/transactions", handlers.TransactionHandler.CreateTransaction)
	app.Get("/transactions", handlers.TransactionHandler.GetAllTransactions)
	app.Get("/transaction/:id", handlers.TransactionHandler.GetTransactionById)
	app.Put("/transaction/:id", handlers.TransactionHandler.UpdateTransaction)

	//Shipment
	app.Post("/shipment", handlers.ShipmentHandler.CreateShipment)
	app.Post("/shipments", handlers.ShipmentHandler.CreateShipment)
	app.Get("/shipments", handlers.ShipmentHandler.GetAllShipments)
	app.Get("/shipment/:id", handlers.ShipmentHandler.GetShipmentByID)
	app.Put("/shipment/:id", handlers.ShipmentHandler.UpdateShipment)

}
