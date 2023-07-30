package repositories

import "dumbsound/models"

type ArtistRepositiry interface {
	CreateArtist(Artist models.Artist) (models.Artist, error)
	FindArtist() ([]models.Artist, error)
	GetArtist(ID int) (models.Artist, error)
	UpdateArtist(Artist models.Artist) (models.Artist, error)
	DeleteArtist(Artist models.Artist, ID int) (models.Artist, error)
}

func (r *repository) CreateArtist(Artist models.Artist) (models.Artist, error) {
	err := r.db.Create(&Artist).Error

	return Artist, err
}

func (r *repository) FindArtist() ([]models.Artist, error) {
	var Artist []models.Artist
	err := r.db.Find(&Artist).Error

	return Artist, err
}

func (r *repository) GetArtist(ID int) (models.Artist, error) {
	var Artist models.Artist
	err := r.db.First(&Artist, "id=?", ID).Error

	return Artist, err
}

func (r *repository) UpdateArtist(Artist models.Artist) (models.Artist, error) {
	err := r.db.Save(&Artist).Error

	return Artist, err
}

func (r *repository) DeleteArtist(Artist models.Artist, ID int) (models.Artist, error) {
	err := r.db.Delete(&Artist).Scan(&Artist).Error

	return Artist, err
}
