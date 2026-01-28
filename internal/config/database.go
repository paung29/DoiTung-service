package config

import (
	"log"
	"time"
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info,   // Log level
			Colorful:      true,          // Color
		},
	)

	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(
		&models.Account{},
		&models.Year{},
		&models.YearFormSetting{},
		&models.Zone{},
		&models.Pole{},
		&models.Cluster{},
		&models.ClusterForm{},
		&models.FlowerForm{},
		&models.PollinationForm{},
		&models.PodForm{},
		&models.PreHarvestForm{},
		&models.HarvestGradingForm{},
		&models.Warehouse{},
		&models.StockMovement{},
	)

	DB = db

}