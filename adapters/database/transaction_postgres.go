package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type TransactionPostgresRepository struct {
	db *gorm.DB
}

func InitiateTransactionPostGresRepository(db *gorm.DB) repositories.TransactionRepository {
	return &TransactionPostgresRepository{
		db: db,
	}
}

func (tpr *TransactionPostgresRepository) CreateTransaction(t entities.Transaction) (entities.Transaction, error) {
	query := "INSERT INTO public.transactions(t_time_stamp, t_net_price, t_image_url) VALUES ($1, $2, $3) RETURNING id, t_time_stamp, t_net_price, t_image_url;"

	var transaction entities.Transaction

	result := tpr.db.Raw(query, t.T_time_stamp, t.T_net_price, t.T_image_url).Scan(&transaction)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	return transaction, nil
}

func (tpr *TransactionPostgresRepository) GetTransactionById(id int) (entities.Transaction, error) {
	var transaction entities.Transaction

	query := "SELECT id, t_time_stamp, t_net_price, t_image_url FROM public.transactions WHERE id = $1;"

	result := tpr.db.Raw(query, id).Scan(&transaction)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	return transaction, nil
}

func (tpr *TransactionPostgresRepository) GetAllTransactions() ([]entities.Transaction, error) {
	query := "SELECT id, t_time_stamp, t_net_price, t_image_url FROM public.transactions"
	var transactions []entities.Transaction

	result := tpr.db.Raw(query).Scan(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (tpr *TransactionPostgresRepository) UpdateTransaction(id int, t entities.Transaction) (entities.Transaction, error) {
	query := "UPDATE public.transactions SET t_time_stamp=$2, t_net_price=$3, t_image_url=$4 WHERE id = $1 RETURNING id, t_time_stamp, t_net_price, t_image_url;"
	var transaction entities.Transaction

	result := tpr.db.Raw(query, id, t.T_time_stamp, t.T_net_price, t.T_image_url).Scan(&transaction)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	return transaction, nil
}
