package pod

type PodService interface {
	CreateOrUpdatePodForm(form PodFormRequest, userId uint) (PodFormResponse, error)
}
