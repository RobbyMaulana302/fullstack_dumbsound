package models

type Music struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Title     string `json:"title" gorm:"type: varchar(255)"`
	Year      string `json:"year" gorm:"type: year"`
	Thumbnail string `json:"thumbnail" gorm:"type: varchar(255)"`
	Attache   string `json:"attache" gorm:"type: varchar(255)"`
	ArtistID  int    `json:"artis_id" gorm:"type: int;  foreignKey:ArtistID"`
	Artist    Artist `json:"artist" gorm:"foreignKey:ArtistID"`
}

type MusicResponse struct {
	Title     string `json:"title" gorm:"type: varchar(255)"`
	Year      string `json:"year" gorm:"type: year"`
	Thumbnail string `json:"thumbnail" gorm:"type: varchar(255)"`
	Attache   string `json:"attache" gorm:"type: varchar(255)"`
	ArtistID  int    `json:"artis_id" gorm:"type: int;  foreignKey:ArtistID"`
	Artist    Artist `json:"artist" gorm:"foreignKey:ArtistID"`
}

func (MusicResponse) TableName() string {
	return "musics"
}
