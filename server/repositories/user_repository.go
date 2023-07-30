package repositories

import "dumbsound/models"

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(Artist models.User) (models.User, error)
}

func (r *repository) FindUsers() ([]models.User, error) {
	var Users []models.User
	err := r.db.Find(&Users).Error

	return Users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var User models.User
	err := r.db.First(&User).Error

	return User, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}
