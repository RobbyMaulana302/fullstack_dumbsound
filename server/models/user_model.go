package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	ListAs    bool      `json:"listAs" gorm:"type: bool"`
	Role      string    `json:"role" gorm:"type: varchar(255)"`
	CreatedAd time.Time `json:"-"`
	UpdateAd  time.Time `json:"-"`
}

type UserResponse struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	ListAs   bool   `json:"listAs" gorm:"type: bool"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
}

func (UserResponse) TableName() string {
	return "users"
}
