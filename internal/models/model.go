package models

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/types"
	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type Account struct {
	AccountID uint
	Name string
	Email string
	PasswordHash string
	Role enums.Role
	
	types.Timestamp
	
}

type Year struct {
	YearID   uint 
	Year int

	Zone []Zone
	YearFormSetting *YearFormSetting

	types.Timestamp

}

type YearFormSetting struct {
	YearID uint
	Year int

	ClusterActive bool
	FlowerActive bool
	PodActive bool
	PreHarvestActive bool
	HarvestGradingActive bool

	yearRef Year

	types.Timestamp

}

type Zone struct {

	ZoneID uint
	YearID uint
	ZoneNo int
	ZoneName string

	YearRef Year
	Pole []Poles
}

type Poles struct {
	PoleID   uint
	ZoneID uint
	PoleNo int

	ZoneRef Zone
	Clusters []Cluster
}

type Cluster struct {
	ClusterID   uint
	PoleID uint
	ClusterNo int

}

type ClusterForm struct {
	ClusterFormID uint
	ClusterID uint
	RecordedBy uint

	Condition enums.Condition
	RecordedDate time.Time

	types.Timestamp
}

type FlowerForm struct {
	FlowerFormID uint
	ClusterID uint
	RecordedBy uint

	TotalFlowers int
}

type PollinationForm struct {
	PollinationFormID uint
	ClusterID uint
	RecordedBy uint

	NumberPods int
	UnsuccessfulPollination int
	GoodFlowers int
	BadFlowers int
	Condition enums.Condition

	RecordedDate time.Time
	types.Timestamp
}

type PodForm struct {
	PodFormID uint
	ClusterID uint
	RecordedBy uint
	NumberPods int
	LostPod int
	RemaingingPods int
}

type PreHarvestForm struct {
	PreHarvestFormID uint
	ClusterID uint
	RecordedBy uint

	NumberPods int
	LostPod int
	RemovedPods int
	PlantsRemoved int
	Condition enums.Condition
}

type HarvestGradingForm struct {
	HarvestGradingFormID uint
	PoleID uint
	RecordedBy uint

	GradeAPlusCount  int
	GradeAPlusWeight int

	GradeACount  int
	GradeAWeight int

	GradeBCount  int
	GradeBWeight int

	GradeCCount  int
	GradeCWeight int

	GradeDPlusCount  int
	GradeDPlusWeight int

	UndersizedCount int
	UndersizedWeight int

	RecordedDate time.Time
	types.Timestamp
}

type StockLocation struct {
	LocationID uint
	LocationName string
}

type StockMovement struct {
	StockMovementID uint
	YearID uint
	LocationID uint
	RecordedBy uint

	Grade string
	MovementType enums.MovementType
	PricePerGram int
	TotalGrams int
	TotalPods int
	Details string
}