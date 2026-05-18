package warehouse

type CreateWarehouseRequest struct {
	Year          uint   `json:"year" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	ActiveStatus  bool   `json:"active_status"`
}

type CreateWarehouseResponse struct {
	Message string `json:"message"`
}
