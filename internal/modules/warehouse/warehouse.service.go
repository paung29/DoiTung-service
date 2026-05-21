package warehouse

type WarehouseService interface {
	CreateWarehouse(form CreateWarehouseRequest) (CreateWarehouseResponse, error)
	GetAllWarehouses() (GetAllWarehousesResponse, error)
	GetWarehouseById(warehouseId uint) (WarehouseDetail, error)
}
