package preharvest

type PreHarvestService interface {
	CreateOreUpdatePreHarvestForm(form PreHarvestFormRequest, userId uint) (PreHarvestFormResponse, error)
}
