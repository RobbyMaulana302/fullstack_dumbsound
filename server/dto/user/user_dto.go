package userdto

import "time"

type UserResponse struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Email       string    `json:"email" gorm:"type: varchar(255)"`
	Password    string    `json:"password" gorm:"type: varchar(255)"`
	FullName    string    `json:"fullname" gorm:"type: varchar(255)"`
	Gender      string    `json:"gender" gorm:"type: varchar(255)"`
	PhoneNumber string    `json:"phone_number" gorm:"type: varchar(255)"`
	Address     string    `json:"address" gorm:"type: varchar(255)"`
	Status      string    `json:"status" gorm:"type: varchar(255)"`
	CreatedAd   time.Time `json:"-"`
	UpdateAd    time.Time `json:"-"`
}
