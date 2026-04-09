package pollination

type PollinationService interface {
	CreateOrUpdatePollinationForm(form PollinationFormRequest, userId uint) (PollinationFormResponse, error)
}
