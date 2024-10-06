package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type TierListPostgres struct {
	db *gorm.DB
}


func ProvideTierListPostgres(db *gorm.DB) repositories.TierListRepository {
	return &TierListPostgres{
		db: db,
	}

}

func (t *TierListPostgres) GetDiscountPercentByUserID(id int) (entities.TierList, error) {
	panic("unimplemented")
}


func (t *TierListPostgres) UpdateTierByUserID(id int, tier int) error {
	panic("unimplemented")
}