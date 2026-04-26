package harvestgrading

type HarvestGradingService interface {
	CreateOrUpdateHarvestGradingForm(form HarvestGradingFormRequest, userId uint) (HarvestGradingFormResponse, error)
	GetHarvestGradingFormDetailsByPoleID(PoleId uint) (HarvestGradingFormDetails, error)
}
