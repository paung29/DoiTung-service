package repository

import "gorm.io/gorm"


func Create[T any] (db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func FindByID[T any] (db *gorm.DB, id uint) (*T, error) {
	var entity T

	if err := db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	
	return &entity, nil
}

func DeleteByID[T any] (db *gorm.DB, id uint) error {
	var entity T
	return db.Delete(&entity, id).Error
}

func Save[T any] (db *gorm.DB, entity T) error {
	return db.Save(entity).Error
}