package year

type YearCreateForm struct {
	Year int `json:"year" validate:"required"`
}

type YearCreateResponse struct {
	Message string `json:"message"`
}

type YearFormSettingStatusChange struct {
	Year int 
	FormID string
	FormName string
	ActiveStatus string 
}

type YearFormSettingStatusChangeResponse struct {
	Message string `json:"message"`
}