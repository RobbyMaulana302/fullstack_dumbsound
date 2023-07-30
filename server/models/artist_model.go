package models

type Artist struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Old         int    `json:"old" gorm:"type: int"`
	Type        string `json:"type" gorm:"type: varchar(255)"`
	StartCareer string `json:"startCareer" gorm:"type: YEAR"`
}

type ArtistResponse struct {
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Old         int    `json:"old" gorm:"type: int"`
	Type        string `json:"type" gorm:"type: varchar(255)"`
	StartCareer string `json:"startCareer" gorm:"type: date"`
}

func (ArtistResponse) TableName() string {
	return "artists"
}
