package models

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/types"
	"github.com/doitung/DoiTung-service/internal/types/enums"
)

type Account struct {
	AccountID 		uint 		`gorm:"primaryKey"`

	Name 			string	
	Email 			string		`gorm:"uniqueIndex;not null"`
	PasswordHash 	string		`gorm:"not null"`
	Role 			enums.Role	`gorm:"type:varchar(20);not null"`
	
	types.Timestamp
	
}

type Year struct {
	YearID  							uint `gorm:"primaryKey"`

	Year 								int  `gorm:"uniqueIndex;not null"`

	Zones []Zone
	YearFormSetting *YearFormSetting

	types.Timestamp

}

type YearFormSetting struct {
	YearID 					uint `gorm:"primaryKey"`

	ClusterActive 			bool
	FlowerActive 			bool
	PodActive 				bool
	PreHarvestActive 		bool
	HarvestGradingActive 	bool

	YearRef 				Year `gorm:"foreignKey:YearID;references:YearID"`

	types.Timestamp

}

type Zone struct {
	ZoneID 		uint 	`gorm:"primaryKey"`

	YearID 		uint	`gorm:"not null;index;uniqueIndex:ux_year_zone_no,priority:1"`
	ZoneNo 		int		`gorm:"not null;uniqueIndex:ux_year_zone_no,priority:2"`

	ZoneName 	string

	YearRef 	Year	`gorm:"foreignKey:YearID;references:YearID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Poles 		[]Pole	`gorm:"foreignKey:ZoneID;references:ZoneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	types.Timestamp
}

type Pole struct {
	PoleID   	uint	`gorm:"primaryKey"`

	ZoneID 		uint	`gorm:"not null;index;uniqueIndex:ux_zone_pole_no,priority:1"`
	PoleNo 		int		`gorm:"not null;uniqueIndex:ux_zone_pole_no,priority:2"`

	ZoneRef 	Zone	`gorm:"foreignKey:ZoneID;references:ZoneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Clusters 	[]Cluster	`gorm:"foreignKey:PoleID;references:PoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	types.Timestamp
}

type Cluster struct {
	ClusterID   uint 	`gorm:"primaryKey"`

	PoleID 		uint	`gorm:"not null;index;uniqueIndex:ux_pole_cluster_no,priority:1"`
	ClusterNo 	int		`gorm:"not null;uniqueIndex:ux_pole_cluster_no,priority:2"`

	PoleRef 	Pole	`gorm:"foreignKey:PoleID;references:PoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	types.Timestamp
}

type ClusterForm struct {
	ClusterFormID 	uint				`gorm:"primaryKey"`

	YearID 			uint				`gorm:"index;not null"`
	ClusterID 		uint				`gorm:"index;not null"`
	RecordedByID 	uint				`gorm:"index;not null"`

	Condition 		enums.Condition		`gorm:"type:varchar(20);not null"`
	RecordedDate 	time.Time			`gorm:"not null"`

	Cluster    		Cluster
	RecordedBy 		Account

	types.Timestamp
}

type FlowerForm struct {
	FlowerFormID 	uint				`gorm:"primaryKey"`

	YearID 			uint				`gorm:"index;not null"`
	ClusterID 		uint				`gorm:"index;not null"`
	RecordedByID 	uint				`gorm:"index;not null"`	

	TotalFlowers 	int
	RecordedDate 	time.Time

	Cluster Cluster
	RecordedBy Account

	types.Timestamp

}

type PollinationForm struct {
	PollinationFormID 			uint	`gorm:"primaryKey"`

	YearID 						uint	`gorm:"index;not null"`
	ClusterID 					uint	`gorm:"index;not null"`
	RecordedByID 				uint	`gorm:"index;not null"`

	NumberPods 					int
	UnsuccessfulPollination 	int
	GoodFlowers 				int
	BadFlowers 					int
	Condition 					enums.Condition	`gorm:"type:varchar(20)"`

	Cluster 					Cluster
	RecordedBy 					Account

	RecordedDate 				time.Time

	types.Timestamp
}

type PodForm struct {
	PodFormID 			uint	`gorm:"primaryKey"`

	YearID   			uint 	`gorm:"index;not null"`
	ClusterID 			uint	`gorm:"index;not null"`
	RecordedByID 		uint	`gorm:"index;not null"`

	NumberPods 			int
	LostPods 			int
	RemainingPods 		int

	Cluster 			Cluster
	RecordedBy 			Account

	RecordedDate  		time.Time

	types.Timestamp
}

type PreHarvestForm struct {
	PreHarvestFormID 	uint	`gorm:"primaryKey"`

	YearID 				uint	`gorm:"index;not null"`
	ClusterID 			uint	`gorm:"index;not null"`
	RecordedByID 		uint	`gorm:"index;not null"`

	NumberPods 			int
	LostPods 			int
	RemovedPods 		int
	PlantsRemoved 		int
	Condition 			enums.Condition	`gorm:"type:varchar(20)"`

	RecordedDate time.Time

	Cluster 			Cluster
	RecordedBy 			Account

	types.Timestamp
}

type HarvestGradingForm struct {
	HarvestGradingFormID uint	`gorm:"primaryKey"`

	YearID 				 uint	`gorm:"index;not null"`
	PoleID 				 uint	`gorm:"index;not null"`
	RecordedByID 		 uint	`gorm:"index;not null"`

	GradeAPlusCount  	int
	GradeAPlusWeight 	int

	GradeACount  		int
	GradeAWeight 		int

	GradeBCount  		int
	GradeBWeight 		int

	GradeCCount  		int
	GradeCWeight 		int

	GradeDCount 	int
	GradeDWeight 	int

	GradeDPlusCount 	int
	GradeDPlusWeight 	int

	UndersizedCount 	int
	UndersizedWeight 	int

	RecordedDate 		time.Time

	Pole 				Pole
	RecordedBy 			Account

	types.Timestamp
}

type Warehouse struct {
	WarehouseID 	uint `gorm:"primaryKey"`

	YearID        	uint   `gorm:"index;not null"`
	WarehouseName 	string `gorm:"not null"`

	types.Timestamp
}

type StockMovement struct {
	StockMovementID 	uint	`gorm:"primaryKey"`

	YearID 				uint	`gorm:"index;not null"`
	RecordedByID 		uint	`gorm:"index;not null"`

	Grade 				enums.Grade			`gorm:"type:varchar(20)"`
	MovementType 		enums.MovementType	`gorm:"type:varchar(20);not null"`
	PricePerGram 		int
	TotalGrams 			int
	TotalPods 			int
	Details 			string

	FromWarehouseID 	*uint
	ToWarehouseID   	*uint

	FromWarehouse 		*Warehouse `gorm:"foreignKey:FromWarehouseID"`
	ToWarehouse   		*Warehouse `gorm:"foreignKey:ToWarehouseID"`

	RecordedBy 			Account

	RecordedDate 		time.Time `gorm:"not null"`
	types.Timestamp
}