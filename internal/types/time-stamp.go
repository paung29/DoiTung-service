package types

import "time"

type Timestamp struct {
	CreatedAt time.Time		`gorm:"autoCreateTime"`
	UpdatedAt time.Time		`gorm:"autoUpdateTime"`
}