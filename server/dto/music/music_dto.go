package musicdto

type MusicRequest struct {
	Title     string `json:"title" form:"title" validate:"required"`
	Year      string `json:"year" form:"year" validate:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Attache   string `json:"attache" form:"attache" validate:"required"`
	ArtistID  int    `json:"artis_id" form:"artist_id" validate:"required"`
}
