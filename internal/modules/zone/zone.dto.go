package zone

type CreateZoneRequest struct {
	Year uint `json:"year" validate:"required,min=4"`
	Name string `json:"name" validate:"required,min=6"`
}

type CreateZoneResponse struct {
	Message string `json:"message"`
}

