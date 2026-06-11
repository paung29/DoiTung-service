package stock

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type CreateCarryOverStockRequest struct {
	Year           uint        `json:"year" validate:"required"`
	ProductionYear *uint       `json:"production_year" validate:"required"`
	WarehouseID    *uint       `json:"warehouse_id" validate:"required"`
	Grade          enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
	TotalGrams     *int        `json:"total_grams" validate:"required,gt=0"`
	TotalPods      *int        `json:"total_pods" validate:"required,gt=0"`
	Details        *string     `json:"details"`
	RecordedDate   time.Time   `json:"recorded_date" validate:"required"`
}

type StockMovementResponse struct {
	Message string `json:"message"`
}

type CreateIncomingStockRequest struct {
	Year           uint        `json:"year" validate:"required"`
	ProductionYear *uint       `json:"production_year" validate:"required"`
	WarehouseID    *uint       `json:"warehouse_id" validate:"required"`
	Grade          enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
	TotalGrams     *int        `json:"total_grams" validate:"required,gt=0"`
	TotalPods      *int        `json:"total_pods" validate:"required,gt=0"`
	Details        *string     `json:"details"`
	RecordedDate   time.Time   `json:"recorded_date" validate:"required"`
}

type CreateIssuedStockRequest struct {
	Year           uint        `json:"year" validate:"required"`
	ProductionYear *uint       `json:"production_year" validate:"required"`
	WarehouseID    *uint       `json:"warehouse_id" validate:"required"`
	CustomerID     *uint       `json:"customer_id" validate:"required"`
	Grade          enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
	PricePerGram   int         `json:"price_per_gram" validate:"required"`
	TotalGrams     int         `json:"total_grams" validate:"required,gt=0"`
	TotalPods      int         `json:"total_pods" validate:"required,gt=0"`
	Details        *string     `json:"details"`
	RecordedDate   time.Time   `json:"recorded_date" validate:"required"`
}

// type UpdateStockMovementRequest struct {
// 	StockMovementID uint        `json:"stock_movement_id" validate:"required"`
// 	ProductionYear  *uint       `json:"production_year" validate:"required"`
// 	WarehouseID     *uint       `json:"warehouse_id" validate:"required"`
// 	CustomerID      *uint       `json:"customer_id"`
// 	Grade           enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
// 	PricePerGram    *int        `json:"price_per_gram"`
// 	TotalGrams      *int        `json:"total_grams" validate:"required,gt=0"`
// 	TotalPods       *int        `json:"total_pods" validate:"required,gt=0"`
// 	Details         *string     `json:"details"`
// }

type StockMovementDetails struct {
	No              uint               `json:"no,omitempty"`
	StockMovementID uint               `json:"stock_movement_id"`
	Date            time.Time          `json:"date"`
	Category        enums.MovementType `json:"category"`
	Grade           enums.Grade        `json:"grade"`
	ProductionYear  int                `json:"production_year"`
	Warehouse       string             `json:"warehouse"`
	TotalGrams      int                `json:"total_grams"`
	TotalPods       int                `json:"total_pods"`
	Details         *string            `json:"details"`
}

type GetAllStockMovementsByYearResponse struct {
	StockMovements []StockMovementDetails `json:"stock_movements"`
}
