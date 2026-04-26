package flower

type FlowerService interface {
	CreateOrUpdateFlowerForm(form FlowerFormRequest, userId uint) (FlowerFormResponse, error)
	GetFlowerFormDetailsByClusterID(clusterId uint) (FlowerFormDetails, error)
}
