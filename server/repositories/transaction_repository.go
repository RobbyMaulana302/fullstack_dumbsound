package repositories

import (
	"dumbsound/models"
	"fmt"
)

type TransactionRepositiry interface {
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetUserTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(status string, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func (r *repository) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&Transaction).Preload("User").Order("id Desc").First(&Transaction).Error

	return Transaction, err
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "id=?", ID).Error

	return transaction, err
}

func (r *repository) GetUserTransaction(ID int) (models.Transaction, error) {
	fmt.Println(ID)
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "user_id=?", ID).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction

	r.db.First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.First(&user, transaction.UserID)
		user.ListAs = true
		r.db.Save(&user)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Scan(&transaction).Error

	return transaction, err
}
