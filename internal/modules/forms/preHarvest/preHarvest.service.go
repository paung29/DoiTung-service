package preharvest

type PreHarvestService interface {
	CreateOrUpdatePreHarvestForm(form PreHarvestFormRequest, userId uint) (PreHarvestFormResponse, error)
}
