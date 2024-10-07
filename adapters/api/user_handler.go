package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type UserHandler struct {
	UserUsecase usecases.UserUseCase
}

func ProvideUserHandler(userUsecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.New(err.Error())
	}

	user, err := uh.UserUsecase.GetUserByID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(user)

}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAllUsers()
	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(users)
}
