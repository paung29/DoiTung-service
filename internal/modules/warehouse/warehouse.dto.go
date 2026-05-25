package warehouse

type CreateWarehouseRequest struct {
	WarehouseName string `json:"warehouse_name" validate:"required"`
	ActiveStatus  bool   `json:"active_status"`
}

type CreateWarehouseResponse struct {
	Message string `json:"message"`
}

type GetAllWarehousesResponse struct {
	Warehouses []WarehouseDetail `json:"warehouses"`
}

type WarehouseDetail struct {
	WarehouseId   uint   `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	ActiveStatus  bool   `json:"active_status"`
}

type UpdateWarehouseRequest struct {
	WarehouseId   uint   `json:"warehouse_id" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	ActiveStatus  bool   `json:"active_status"`
}

type UpdateWarehouseResponse struct {
	Message string `json:"message"`
}
