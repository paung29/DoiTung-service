package stock

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type CreateCarryOverRequest struct {
	YearID           uint        `json:"year_id" validate:"required"`
	ProductionYearID *uint       `json:"production_year_id" validate:"required"`
	WarehouseID      *uint       `json:"warehouse_id" validate:"required"`
	Grade            enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
	TotalGrams       *int        `json:"total_grams" validate:"required"`
	TotalPods        *int        `json:"total_pods" validate:"required"`
	Details          *string     `json:"details"`
	RecordedDate     time.Time   `json:"recorded_date" validate:"required"`
}

type StockMovementResponse struct {
	Message string `json:"message"`
}

type CreateIncomingStockRequest struct {
	YearID           uint        `json:"year_id" validate:"required"`
	ProductionYearID *uint       `json:"production_year_id" validate:"required"`
	WarehouseID      *uint       `json:"warehouse_id" validate:"required"`
	Grade            enums.Grade `json:"grade" validate:"required,oneof=A_PLUS A B C D D_PLUS"`
	TotalGrams       *int        `json:"total_grams" validate:"required"`
	TotalPods        *int        `json:"total_pods" validate:"required"`
	Details          *string     `json:"details"`
	RecordedDate     time.Time   `json:"recorded_date" validate:"required"`
}
