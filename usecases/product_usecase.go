package usecases

import (
	"errors"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type ProductUsecase interface {
	CreateProduct(p entities.Product) error
	GetProductByID(id int) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
}

type ProductService struct {
	repo repositories.ProductRepository
}

func InitiateProductsService(repo repositories.ProductRepository) ProductUsecase {
	return &ProductService{
		repo: repo,
	}
}

func (ps *ProductService) CreateProduct(p entities.Product) error {

	if err := ps.repo.CreateProduct(p); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (ps *ProductService) GetProductByID(id int) (entities.Product, error) {
	p, err := ps.repo.GetProductByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	return p, nil
}

func (ps *ProductService) GetAllProducts() ([]entities.Product, error) {
	p_list, err := ps.repo.GetAllProducts()

	if err != nil {
		return []entities.Product{}, err
	}

	return p_list, nil
}
