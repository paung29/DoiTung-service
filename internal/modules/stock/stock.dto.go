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
	TotalGrams     *float64    `json:"total_grams" validate:"required,gt=0"`
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
	TotalGrams     *float64    `json:"total_grams" validate:"required,gt=0"`
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
	TotalGrams     float64     `json:"total_grams" validate:"required,gt=0"`
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
	Date            string             `json:"date"`
	Category        enums.MovementType `json:"category"`
	Grade           enums.Grade        `json:"grade"`
	ProductionYear  int                `json:"production_year"`
	Warehouse       string             `json:"warehouse"`
	TotalGrams      float64            `json:"total_grams"`
	TotalPods       int                `json:"total_pods"`
	Details         *string            `json:"details"`
}

type GetAllStockMovementsByYearResponse struct {
	StockMovements []StockMovementDetails `json:"stock_movements"`
}

type CustomerStockTableByYearResponse struct {
	CustomerStockTable []CustomerStockTableItem `json:"customer_stock_table"`
}

type CustomerStockTableItem struct {
	CustomerID   int     `json:"customer_id"`
	No           int     `json:"no"`
	CustomerName string  `json:"customer_name"`
	GradeA       float64 `json:"grade_a"`
	GradeB       float64 `json:"grade_b"`
	GradeC       float64 `json:"grade_c"`
	GradeFailed  float64 `json:"grade_failed"`
	TotalWeight  float64 `json:"total_weight"`
	Note         *string `json:"note"`
}

type StockOverviewResponse struct {
	TotalPodInStock   int                  `json:"total_pod_in_stock"`
	TotalGramInStock  float64              `json:"total_gram_in_stock"`
	TotalKgInStock    float64              `json:"total_kg_in_stock"`
	IncomingStockPod  int                  `json:"incoming_stock_pod"`
	IncomingStockGram float64              `json:"incoming_stock_gram"`
	IncomingStockKg   float64              `json:"incoming_stock_kg"`
	IssuedStockPod    int                  `json:"issued_stock_pod"`
	IssuedStockGram   float64              `json:"issued_stock_gram"`
	IssuedStockKg     float64              `json:"issued_stock_kg"`
	GradeSummary      []GradeSummaryItem   `json:"grade_summary"`
	MonthlySummary    []MonthlySummaryItem `json:"monthly_summary"`
}

type GradeSummaryItem struct {
	Grade      enums.Grade `json:"grade"`
	TotalPod   int         `json:"total_pod"`
	TotalGram  float64     `json:"total_gram"`
	TotalKg    float64     `json:"total_kg"`
	Percentage float64     `json:"percentage"`
}

type MonthlySummaryItem struct {
	Month          int     `json:"month"`
	MonthName      string  `json:"month_name"`
	StockInWeight  float64 `json:"stock_in_weight"`
	StockOutWeight float64 `json:"stock_out_weight"`
	TotalWeight    float64 `json:"total_weight"`
}
