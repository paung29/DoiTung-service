package flower

type FlowerService interface {
	CreateFlowerForm(form FlowerFormRequest, userId uint) (FlowerFormResponse, error)
}
