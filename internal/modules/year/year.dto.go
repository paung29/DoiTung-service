package year

type YearCreateForm struct {
	Year int `json:"year" validate:"required"`
}

type YearCreateResponse struct {
	Message string `json:"message"`
}

type YearFormSettingStatusChange struct {
	Year int `json:"year" validate:"required"`
	FormName string	`json:"formName" validate:"required,oneof=cluster flower pollination pod preHarvest harvestGrading"`
	ActiveStatus *bool 	`json:"activeStatus" validate:"required"`
}

type YearFormSettingStatusChangeResponse struct {
	Message string `json:"message"`
}

type GetYearResponse struct {
	Years []string `json:"years"`
}