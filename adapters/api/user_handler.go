package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/domain/response"
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

	if user == nil {
		return c.JSON(fiber.Map{
			"message": "can't find user",
		})
	}

	if err != nil {
		return errors.New(err.Error())
	}

	response := response.GetUserResponse{
		ID:           user.ID,
		CredentialID: user.CredentialID,
		Name:         user.Name,
		Email:        user.Email,
		Status:       user.Status,
		Role:         user.Role,
		TierRank:     user.TierRank,
	}

	return c.JSON(response)

}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAllUsers()
	
	if users == nil {
		return c.JSON(fiber.Map{
			"message": "can't find user",
		})
	}

	if err != nil {
		return errors.New(err.Error())
	}

	var userResponses []response.GetUserResponse

	for _, user := range *users {
		userResponse := response.GetUserResponse{
			ID:           user.ID,
			CredentialID: user.CredentialID,
			Name:         user.Name,
			PhoneNumber:  user.PhoneNumber,
			Email:        user.Email,
			Status:       user.Status,
			Role:         user.Role,
			TierRank:     user.TierRank,
		}
		userResponses = append(userResponses, userResponse)
	}

	result := response.GetUsersResponse{
		Users: userResponses,
	}

	return c.JSON(result)
}

func (uh *UserHandler) UpdateTierByUserID(c *fiber.Ctx) error {

	req := new(request.UpdateTierByUserIDRequest)

	if err := c.BodyParser(&req); err != nil {
		return errors.New(err.Error())
	}

	user, err := uh.UserUsecase.UpdateTierByUserID(req)

	// if user == nil {
	// 	return c.JSON(fiber.Map{
	// 		"message": "can't find user",
	// 	})
	// }

	if err != nil {
		return errors.New(err.Error())
	}

	result := response.UserResponse{
		ID:           user.ID,
		CredentialID: user.CredentialID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		Status:       user.Status,
		Role:         user.Role,
		TierRank:     user.TierRank,
	}

	return c.JSON(result)

}
