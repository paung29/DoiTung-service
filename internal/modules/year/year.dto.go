package year

type YearCreateForm struct {
	Year int `json:"year" validate:"required"`
}

type YearCreateResponse struct {
	Message string `json:"message"`
}

type YearFormSettingStatusChange struct {
	Year         int    `json:"year" validate:"required"`
	FormName     string `json:"formName" validate:"required,oneof=cluster flower pollination pod preHarvest harvestGrading"`
	ActiveStatus *bool  `json:"activeStatus" validate:"required"`
}

type YearFormSettingStatusChangeResponse struct {
	Message string `json:"message"`
}

type GetYearResponse struct {
	Years []string `json:"years"`
}

type GetYearDetailsLists struct {
	YearDetails []YearDetails `json:"yearDetails"`
}

type YearDetails struct {
	YearId               uint `json:"yearId"`
	Year                 int  `json:"year"`
	ClusterActive        bool `json:"clusterActive"`
	FlowerActive         bool `json:"flowerActive"`
	PollinationActive    bool `json:"pollinationActive"`
	PodActive            bool `json:"podActive"`
	PreHarvestActive     bool `json:"preHarvestActive"`
	HarvestGradingActive bool `json:"harvestGradingActive"`
}
