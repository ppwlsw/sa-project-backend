package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func ProvideUserPostgresRepository(db *gorm.DB) repositories.UserRepository {
	return &UserPostgresRepository{db: db}
}

func (upr *UserPostgresRepository) CreateUser(newUser *entities.User) error {
	query := "INSERT INTO public.users(credential_id, name, phone_number, email, password, status, role, tier_rank) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	err := upr.db.Exec(query, newUser.CredentialID, newUser.Name, newUser.PhoneNumber, newUser.Email, newUser.Password, newUser.Status, newUser.Role, newUser.TierRank)
	if err != nil {
		return err.Error
	}

	return nil
}

func (upr *UserPostgresRepository) GetUserByID(id int) (*entities.User, error) {
	var user *entities.User

	query := "SELECT credential_id, name, phone_number, email, password, status, role, tier_rank FROM users WHERE id = $1"

	result := upr.db.Raw(query, id).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (upr *UserPostgresRepository) GetAllUsers() (*[]entities.User, error) {
	query := "SELECT credential_id, name, phone_number, email, password, status, role, tier_rank FROM users"
	var users *[]entities.User

	result := upr.db.Raw(query).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil

}

func (upr *UserPostgresRepository) FindUserByEmail(email string) (*entities.User, error) {
	query := "SELECT credential_id, name, phone_number, email, password, status, role, tier_rank FROM users WHERE email = $1"
	var user *entities.User

	result := upr.db.Raw(query, email).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
