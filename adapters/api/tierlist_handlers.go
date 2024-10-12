package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type TierListHandler struct {
	TierListUsecase usecases.TierListUsecase
}

func InitiateTierListHandler(TierListUsecase usecases.TierListUsecase) *TierListHandler {
	return &TierListHandler{
		TierListUsecase: TierListUsecase,
	}
}

func (tlh *TierListHandler) GetDiscountPercentByUserID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	discount, err := tlh.TierListUsecase.GetDiscountPercentByUserID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(discount)
}


