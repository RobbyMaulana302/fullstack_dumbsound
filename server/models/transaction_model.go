package models

type Transaction struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string `json:"start_date" gorm:"type: date"`
	DueDate   string `json:"due_date" gorm:"type: date"`
	Status    string `json:"status" gorm:"type: varchar(255)"`
	Price     int    `json:"price" gorm:"type: int"`
	UserID    int    `json:"user_id" gorm:"type: int(20); foreignKey:UserID"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
}

type TransactionResponse struct {
	StartDate string `json:"start_date" gorm:"type: date"`
	DueDate   string `json:"due_date" gorm:"type: date"`
	Status    string `json:"status" gorm:"type: varchar(255)"`
	Price     int    `json:"price" gorm:"type: int"`
	UserID    int    `json:"user_id" gorm:"type: int(20); foreignKey:UserID"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
