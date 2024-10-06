package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type AuthHandler struct {
	AuthUsecase usecases.AuthUsecase
}

func ProvideAuthHandler(usecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		AuthUsecase: usecase,
	}
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can not Parse Body",
		})
	}

	if err := ah.AuthUsecase.Register(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't Create User",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create Successfully",
	})

}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	err := ah.AuthUsecase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "login successful",
	})

}
