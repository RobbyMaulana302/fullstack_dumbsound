package repositories

import "dumbsound/models"

type MusicRepositiry interface {
	CreateMusic(music models.Music) (models.Music, error)
	FindMusic() ([]models.Music, error)
	GetMusic(ID int) (models.Music, error)
}

func (r *repository) CreateMusic(music models.Music) (models.Music, error) {
	err := r.db.Create(&music).Preload("Artist").Order("id Desc").First(&music).Error

	return music, err
}

func (r *repository) FindMusic() ([]models.Music, error) {
	var music []models.Music
	err := r.db.Preload("Artist").Find(&music).Error

	return music, err
}

func (r *repository) GetMusic(ID int) (models.Music, error) {
	var music models.Music
	err := r.db.Preload("Artist").First(&music, "id=?", ID).Error

	return music, err
}
