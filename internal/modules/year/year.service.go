package year

import "gorm.io/gorm"


type YearService interface {
	CreateYear(form YearCreateForm) (YearCreateResponse, error)
}

type service struct {
	db *gorm.DB
	yearRepo YearRepository
}

func NewYearService(db *gorm.DB ,yearRepo YearRepository) YearService {
	return &service{
		db : db,
		yearRepo: yearRepo,
	}
}