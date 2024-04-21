package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type storage struct {
	DB *gorm.DB
}

func NewStorage() *storage {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading: %v", err)
	}
	dbPort := 5432
	hostName := "127.0.0.1"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		hostName,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&entities.Expression{}, &entities.User{})

	return &storage{
		DB: db,
	}
}
