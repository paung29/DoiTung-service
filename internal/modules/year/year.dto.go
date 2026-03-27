package year

type YearCreateForm struct {
	Year int `json:"year" validate:"required"`
}

type YearCreateResponse struct {
	Message string `json:"message"`
}