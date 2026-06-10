package year

import "gorm.io/gorm"

type YearService interface {
	CreateYear(form YearCreateForm) (YearCreateResponse, error)
	ChangeYearFormSettingStatus(form YearFormSettingStatusChange) (YearFormSettingStatusChangeResponse, error)
	GetYear() (GetYearResponse, error)
	GetYearDetails(year int) (YearSettingDetailsResponse, error)
	GetYearManagementTable() (YearManagementListResponse, error)
}

type service struct {
	db       *gorm.DB
	yearRepo YearRepository
}

func NewYearService(db *gorm.DB, yearRepo YearRepository) YearService {
	return &service{
		db:       db,
		yearRepo: yearRepo,
	}
}
