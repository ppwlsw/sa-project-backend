package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type ProductHandler struct {
	ProductUsecase usecases.ProductUsecase
}

func InitiateProductHandler(productUsecase usecases.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var newProduct entities.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return errors.New(err.Error())
	}

	if err := ph.ProductUsecase.CreateProduct(newProduct); err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Create Product Succesful",
	})
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	product, err := ph.ProductUsecase.GetProductByID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(product)
}

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := ph.ProductUsecase.GetAllProducts()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(products)
}