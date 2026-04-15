package harvestgrading

type HarvestGradingService interface {
	CreateOrUpdateHarvestGradingForm(form HarvestGradingFormRequest, userId uint) (HarvestGradingFormResponse, error)
}
