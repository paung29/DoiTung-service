package harvestgrading

type HarvestGradingFormRequest struct {
	PoleId           uint     `json:"poleId" validate:"required,number"`
	GradeAPlusCount  *uint    `json:"gradeAPlusCount" validate:"required,number"`
	GradeAPlusWeight *float64 `json:"gradeAPlusWeight" validate:"required,number"`
	GradeACount      *uint    `json:"gradeACount" validate:"required,number"`
	GradeAWeight     *float64 `json:"gradeAWeight" validate:"required,number"`
	GradeBCount      *uint    `json:"gradeBCount" validate:"required,number"`
	GradeBWeight     *float64 `json:"gradeBWeight" validate:"required,number"`
	GradeCCount      *uint    `json:"gradeCCount" validate:"required,number"`
	GradeCWeight     *float64 `json:"gradeCWeight" validate:"required,number"`
	GradeDPlusCount  *uint    `json:"gradeDPlusCount" validate:"required,number"`
	GradeDPlusWeight *float64 `json:"gradeDPlusWeight" validate:"required,number"`
	UndersizedCount  *uint    `json:"undersizedCount" validate:"required,number"`
	UndersizedWeight *float64 `json:"undersizedWeight" validate:"required,number"`
	RottenCount      *uint    `json:"rottenCount" validate:"required,number"`
	RottenWeight     *float64 `json:"rottenWeight" validate:"required,number"`
}

type HarvestGradingFormResponse struct {
	Message string `json:"message"`
}

type HarvestGradingFormDetails struct {
	No                     uint    `json:"no,omitempty"`
	PoleId                 uint    `json:"poleId"`
	Year                   uint    `json:"year"`
	Location               string  `json:"location"`
	PoleNo                 uint    `json:"poleNo"`
	GradeAPlusCount        uint    `json:"gradeAPlusCount"`
	GradeAPlusWeight       float64 `json:"gradeAPlusWeight"`
	GradeACount            uint    `json:"gradeACount"`
	GradeAWeight           float64 `json:"gradeAWeight"`
	GradeBCount            uint    `json:"gradeBCount"`
	GradeBWeight           float64 `json:"gradeBWeight"`
	GradeCCount            uint    `json:"gradeCCount"`
	GradeCWeight           float64 `json:"gradeCWeight"`
	GradeDPlusCount        uint    `json:"gradeDPlusCount"`
	GradeDPlusWeight       float64 `json:"gradeDPlusWeight"`
	UndersizedCount        uint    `json:"undersizedCount"`
	UndersizedWeight       float64 `json:"undersizedWeight"`
	RottenCount            uint    `json:"rottenCount"`
	RottenWeight           float64 `json:"rottenWeight"`
	HarvestGradingFormDone bool    `json:"harvestGradingFormDone"`
	RecordedBy             string  `json:"recordedBy,omitempty"`
	Date                   string  `json:"date,omitempty"`
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

type HarvestGradingFormLists struct {
	HarvestGradingForms []HarvestGradingFormDetails `json:"harvestGradingForms"`
}
