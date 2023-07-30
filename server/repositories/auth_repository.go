package repositories

import "dumbsound/models"

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(Email string) (models.User, error)
	CheckAuth(ID int, userRole string) (models.User, error)
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email=?", email).First(&user).Error

	return user, err
}

func (r *repository) CheckAuth(ID int, userRole string) (models.User, error) {
	var user models.User
	err := r.db.Where("role=?", userRole).First(&user, ID).Error

	return user, err
}
