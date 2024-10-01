package usecases

import (
	"errors"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type UserUseCase interface {
	CreateUser(user entities.User) error
	GetUserByID(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
}

type UserService struct {
	repo repositories.UserRepository
}



func ProvideUserService(repo repositories.UserRepository) UserUseCase {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) CreateUser(user entities.User) error {
	
	if err := us.repo.CreateUser(user); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (us *UserService) GetUserByID(id int) (entities.User, error) {
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (us *UserService) GetAllUsers() ([]entities.User, error) {
	users , err := us.repo.GetAllUsers()

	if err != nil {
		return []entities.User{}, err
	}

	return users,nil
}