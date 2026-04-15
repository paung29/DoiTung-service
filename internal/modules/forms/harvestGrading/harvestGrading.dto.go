package harvestgrading

type HarvestGradingFormRequest struct {
	Year             uint  `json:"year" validate:"required,number"`
	ZoneNo           uint  `json:"zone-no" validate:"required,number"`
	PoleNo           uint  `json:"pole-no" validate:"required,number"`
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
