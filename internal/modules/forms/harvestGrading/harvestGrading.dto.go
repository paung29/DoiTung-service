package harvestgrading

type HarvestGradingFormRequest struct {
	PoleId           uint  `json:"poleId" validate:"required,number"`
	GradeAPlusCount  *uint `json:"gradeAPlusCount" validate:"required,number"`
	GradeAPlusWeight *uint `json:"gradeAPlusWeight" validate:"required,number"`
	GradeACount      *uint `json:"gradeACount" validate:"required,number"`
	GradeAWeight     *uint `json:"gradeAWeight" validate:"required,number"`
	GradeBCount      *uint `json:"gradeBCount" validate:"required,number"`
	GradeBWeight     *uint `json:"gradeBWeight" validate:"required,number"`
	GradeCCount      *uint `json:"gradeCCount" validate:"required,number"`
	GradeCWeight     *uint `json:"gradeCWeight" validate:"required,number"`
	GradeDPlusCount  *uint `json:"gradeDPlusCount" validate:"required,number"`
	GradeDPlusWeight *uint `json:"gradeDPlusWeight" validate:"required,number"`
	UndersizedCount  *uint `json:"undersizedCount" validate:"required,number"`
	UndersizedWeight *uint `json:"undersizedWeight" validate:"required,number"`
}

type HarvestGradingFormResponse struct {
	Message string `json:"message"`
}

type HarvestGradingFormDetails struct {
	PoleId                 uint   `json:"poleId"`
	Year                   uint   `json:"year"`
	Location               string `json:"location"`
	PoleNo                 uint   `json:"poleNo"`
	GradeAPlusCount        uint   `json:"gradeAPlusCount"`
	GradeAPlusWeight       uint   `json:"gradeAPlusWeight"`
	GradeACount            uint   `json:"gradeACount"`
	GradeAWeight           uint   `json:"gradeAWeight"`
	GradeBCount            uint   `json:"gradeBCount"`
	GradeBWeight           uint   `json:"gradeBWeight"`
	GradeCCount            uint   `json:"gradeCCount"`
	GradeCWeight           uint   `json:"gradeCWeight"`
	GradeDPlusCount        uint   `json:"gradeDPlusCount"`
	GradeDPlusWeight       uint   `json:"gradeDPlusWeight"`
	UndersizedCount        uint   `json:"undersizedCount"`
	UndersizedWeight       uint   `json:"undersizedWeight"`
	HarvestGradingFormDone bool   `json:"harvestGradingFormDone"`
}

type HarvestGradingFormHistoriesResponse struct {
	HarvestGradingFormHistories []HarvestGradingFormHistory `json:"harvestGradingFormHistories"`
}

type HarvestGradingFormHistory struct {
	PoleId                 uint   `json:"poleId"`
	Location               string `json:"location"`
	PoleNo                 uint   `json:"poleNo"`
	HarvestGradingFormDone bool   `json:"harvestGradingFormDone"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
}
