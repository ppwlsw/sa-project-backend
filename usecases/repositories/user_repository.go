package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() (*[]entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
}
