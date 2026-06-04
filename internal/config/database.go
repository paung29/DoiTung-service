package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	godotenv.Load(".env")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err == nil {
			break
		}

		log.Printf("database connection attempt %d failed: %v", i, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
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
		&models.StockBalance{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	SeedAccounts()
}
