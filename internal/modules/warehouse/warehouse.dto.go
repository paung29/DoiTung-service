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

type WarehouseTableByYearResponse struct {
	TotalWarehouses       int                  `json:"total_warehouses"`
	TotalActiveWarehouses int                  `json:"total_active_warehouses"`
	TotalStocksPods       int                  `json:"total_stocks_pods"`
	TotalStocksWeights    float64              `json:"total_stocks_weights"`
	WarehouseTable        []WarehouseTableItem `json:"warehouse_table"`
}

type WarehouseTableItem struct {
	WarehouseId   uint   `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	ActiveStatus  bool   `json:"active_status"`

	TotalPods    int     `json:"total_pods"`
	TotalWeights float64 `json:"total_weights"`

	DistributedPods    int     `json:"distributed_pods"`
	DistributedWeights float64 `json:"distributed_weights"`

	RemainingPods    int     `json:"remaining_pods"`
	RemainingWeights float64 `json:"remaining_weights"`
}
