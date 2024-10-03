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
	query := "INSERT INTO public.products(p_name, p_description, p_location, p_amount, p_price, image_url_1, image_url_2, image_url_3) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	err := ppr.db.Exec(query, p.P_name, p.P_description, p.P_location, p.P_amount, p.P_price, p.Image_url_1, p.Image_url_2, p.Image_url_3)
	if err != nil {
		return err.Error
	}

	return nil
}

func (ppr *ProductPostgresRepository) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product

	query := "SELECT id, p_name, p_description, p_location, p_amount, p_price, image_url_1, image_url_2, image_url_3 FROM public.products WHERE id = $1;"

	result := ppr.db.Raw(query, id).Scan(&product)

	if result.Error != nil {
		return entities.Product{}, result.Error
	}

	return product, nil
}

func (ppr *ProductPostgresRepository) GetAllProducts() ([]entities.Product, error) {
	query := "SELECT id, p_name, p_description, p_location, p_amount, p_price, image_url_1, image_url_2, image_url_3 FROM public.products;"
	var products []entities.Product

	result := ppr.db.Raw(query).Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
