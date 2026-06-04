package preharvest

type PreHarvestService interface {
	CreateOrUpdatePreHarvestForm(form PreHarvestFormRequest, userId uint) (PreHarvestFormResponse, error)
	GetPreHarvestFormDetails(clusterId uint) (PreHarvestFormDetails, error)
	GetPreHarvestFormHistories(userId uint, year uint) (PreHarvestFormHistoriesResponse, error)
	GetPreHarvestFormByZoneId(zoneId uint) (PreHarvestFormLists, error)
}
