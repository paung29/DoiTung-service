package warehouse

type CreateWarehouseRequest struct {
	WarehouseName string `json:"warehouse_name" validate:"required"`
	ActiveStatus  bool   `json:"active_status"`
}

type CreateWarehouseResponse struct {
	Message string `json:"message"`
}
