package pod

type PodService interface {
	CreateOrUpdatePodForm(form PodFormRequest, userId uint) (PodFormResponse, error)
	GetPodFormDetails(clusterId uint) (PodFormDetails, error)
	GetPodFormHistories(userId uint, year uint) (PodFormHistoriesResponse, error)
	GetPodFormsByZoneId(zoneId uint) (PodFormLists, error)
}
