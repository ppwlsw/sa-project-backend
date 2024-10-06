package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type TierListRepository interface {
	GetDiscountPercentByUserID(id int) (entities.TierList, error)
	UpdateTierByUserID(id int, tier int) error 
}