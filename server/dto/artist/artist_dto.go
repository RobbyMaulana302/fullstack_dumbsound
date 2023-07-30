package artistdto

type ArtistRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Old         string `json:"old" form:"old" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required"`
	StartCareer string `json:"start_career" form:"start_career" validate:"required"`
}

type ArtistResponse struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Old         string `json:"old" form:"old" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required"`
	StartCareer string `json:"start_career" form:"type" validate:"required"`
}

type ArtistUpdateRequest struct {
	Name        string `json:"name" form:"name"`
	Old         string `json:"old" form:"old"`
	Type        string `json:"type" form:"type"`
	StartCareer string `json:"start_career" form:"start_career"`
}
