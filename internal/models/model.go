package models

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/types"
	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type Account struct {
	AccountID uint `gorm:"primaryKey"`

	Name         string
	Email        string     `gorm:"uniqueIndex;not null"`
	PasswordHash string     `gorm:"not null"`
	Role         enums.Role `gorm:"type:varchar(20);not null"`

	types.Timestamp
}

type Year struct {
	YearID uint `gorm:"primaryKey"`

	Year int `gorm:"uniqueIndex;not null"`

	Zones            []Zone
	YearFormSetting  *YearFormSetting
	ClusterForms     []ClusterForm
	FlowerForms      []FlowerForm
	PollinationForms []PollinationForm
	PodForms         []PodForm
	PreHarvestForms  []PreHarvestForm
	HarvestForms     []HarvestGradingForm
	Warehouses       []Warehouse
	StockMovements   []StockMovement

	types.Timestamp
}

type YearFormSetting struct {
	YearID uint `gorm:"primaryKey"`

	ClusterActive        bool
	FlowerActive         bool
	PollinationActive    bool
	PodActive            bool
	PreHarvestActive     bool
	HarvestGradingActive bool

	Year Year `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type Zone struct {
	ZoneID uint `gorm:"primaryKey"`

	YearID uint `gorm:"not null;index;uniqueIndex:ux_year_zone_no,priority:1;uniqueIndex:ux_year_zone_name,priority:1"`
	ZoneNo int  `gorm:"not null;uniqueIndex:ux_year_zone_no,priority:2"`

	ZoneName string `gorm:"not null;uniqueIndex:ux_year_zone_name,priority:2"`

	Year Year   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Poles   []Pole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type Pole struct {
	PoleID uint `gorm:"primaryKey"`

	ZoneID uint `gorm:"not null;index;uniqueIndex:ux_zone_pole_no,priority:1"`
	PoleNo int  `gorm:"not null;uniqueIndex:ux_zone_pole_no,priority:2"`

	HarvestGradingFormDone bool `gorm:"default:false"`

	Zone             Zone                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Clusters            []Cluster            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	HarvestGradingForms []HarvestGradingForm `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type Cluster struct {
	ClusterID uint `gorm:"primaryKey"`

	PoleID    uint `gorm:"not null;index;uniqueIndex:ux_pole_cluster_no,priority:1"`
	ClusterNo int  `gorm:"not null;uniqueIndex:ux_pole_cluster_no,priority:2"`

	ClusterFormDone     bool `gorm:"default:false"`
	FlowerFormDone      bool `gorm:"default:false"`
	PodFormDone         bool `gorm:"default:false"`
	PollinationFormDone bool `gorm:"default:false"`
	PreHarvestFormDone  bool `gorm:"default:false"`

	Pole            Pole               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ClusterForms       []ClusterForm      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	FlowerForms        []FlowerForm       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	PollinationForms   []PollinationForm  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	PodForms           []PodForm          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	PreHarvestForms    []PreHarvestForm   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type ClusterForm struct {
	ClusterFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_cluster_form,priority:1"`
	ClusterID    uint `gorm:"not null;uniqueIndex:ux_year_cluster_form,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	Condition    enums.Condition `gorm:"type:varchar(20);not null"`
	RecordedDate time.Time       `gorm:"not null"`

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Cluster    Cluster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type FlowerForm struct {
	FlowerFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_cluster_flower,priority:1"`
	ClusterID    uint `gorm:"not null;uniqueIndex:ux_year_cluster_flower,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	TotalFlowers int
	Condition    enums.Condition `gorm:"type:varchar(20)"`
	Done         bool            `gorm:"default:false"`
	RecordedDate time.Time

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Cluster    Cluster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type PollinationForm struct {
	PollinationFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_cluster_pollination,priority:1"`
	ClusterID    uint `gorm:"not null;uniqueIndex:ux_year_cluster_pollination,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	NumberPods              int
	UnsuccessfulPollination int
	GoodFlowers             int
	BadFlowers              int
	Condition               enums.Condition `gorm:"type:varchar(20)"`
	RecordedDate            time.Time

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Cluster    Cluster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type PodForm struct {
	PodFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_cluster_pod,priority:1"`
	ClusterID    uint `gorm:"not null;uniqueIndex:ux_year_cluster_pod,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	NumberPods    int
	LostPods      int
	RemainingPods int
	RecordedDate  time.Time

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Cluster    Cluster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type PreHarvestForm struct {
	PreHarvestFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_cluster_preharvest,priority:1"`
	ClusterID    uint `gorm:"not null;uniqueIndex:ux_year_cluster_preharvest,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	NumberPods    int
	LostPods      int
	RemovedPods   int
	PlantsRemoved int
	Condition     enums.Condition `gorm:"type:varchar(20)"`
	RecordedDate  time.Time

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Cluster    Cluster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type HarvestGradingForm struct {
	HarvestGradingFormID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"not null;uniqueIndex:ux_year_pole_harvest,priority:1"`
	PoleID       uint `gorm:"not null;uniqueIndex:ux_year_pole_harvest,priority:2"`
	RecordedByID uint `gorm:"index;not null"`

	GradeAPlusCount  int
	GradeAPlusWeight int
	GradeACount      int
	GradeAWeight     int
	GradeBCount      int
	GradeBWeight     int
	GradeCCount      int
	GradeCWeight     int
	GradeDCount      int
	GradeDWeight     int
	GradeDPlusCount  int
	GradeDPlusWeight int
	UndersizedCount  int
	UndersizedWeight int

	RecordedDate time.Time

	Year    Year    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Pole       Pole    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RecordedBy Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type Warehouse struct {
	WarehouseID uint `gorm:"primaryKey"`

	YearID        uint   `gorm:"not null;uniqueIndex:ux_year_warehouse_name,priority:1"`
	WarehouseName string `gorm:"not null;uniqueIndex:ux_year_warehouse_name,priority:2"`

	Year Year `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	types.Timestamp
}

type StockMovement struct {
	StockMovementID uint `gorm:"primaryKey"`

	YearID       uint `gorm:"index;not null"`
	RecordedByID uint `gorm:"index;not null"`

	Grade        enums.Grade        `gorm:"type:varchar(20)"`
	MovementType enums.MovementType `gorm:"type:varchar(20);not null"`
	PricePerGram int
	TotalGrams   int
	TotalPods    int
	Details      string

	FromWarehouseID *uint
	ToWarehouseID   *uint

	Year       Year       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	FromWarehouse *Warehouse `gorm:"foreignKey:FromWarehouseID;references:WarehouseID"`
	ToWarehouse   *Warehouse `gorm:"foreignKey:ToWarehouseID;references:WarehouseID"`
	RecordedBy    Account    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	RecordedDate time.Time `gorm:"not null"`

	types.Timestamp
}