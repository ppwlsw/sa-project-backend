package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	db *gorm.DB
}

func InitiateProductPostGresRepository(db *gorm.DB) repositories.ProductRepository {
	return &ProductPostgresRepository{db: db}
}

func (ppr *ProductPostgresRepository) CreateProduct(p entities.Product) error {
	query := "INSERT INTO public.products(p_id, p_name, p_amount, p_price, image_url_1, image_url_2, image_url_3, p_location)VALUES ($1, $2, $3, $4, $5, $6);"

	err := ppr.db.Exec(query, p.P_name, p.P_amount, p.P_price, p.Image_url_1, p.Image_url_2, p.Image_url_3, p.P_location)
	if err != nil {
		return err.Error
	}

	return nil
}

func (ppr *ProductPostgresRepository) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product

	query := "SELECT p_id, p_name, p_amount, p_price, image_url_1, image_url_2, image_url_3, p_location FROM public.products WHERE p_id = $1;"

	result := ppr.db.Raw(query, id).Scan(&product)

	if result.Error != nil {
		return entities.Product{}, result.Error
	}

	return product, nil
}

func (ppr *ProductPostgresRepository) GetAllProducts() ([]entities.Product, error) {
	query := "SELECT p_id, p_name, p_amount, p_price, image_url_1, image_url_2, image_url_3, p_location FROM public.products;"
	var products []entities.Product

	result := ppr.db.Raw(query).Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
