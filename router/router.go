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

	handlers := api.ProvideHandlers(userHandler, productHandler)

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
	app.Get("products/:id", handlers.ProductHandler.GetUserByID)

}
