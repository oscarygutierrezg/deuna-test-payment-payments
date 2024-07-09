package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"payment-payments-api/internal/models"
)

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
		return nil, err
	}
	log.Println("Conexión exitosa a PostgreSQL")
	return DB, nil
}

func Migrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Println("Error al migrar la base de datos:", err)
		return err
	}
	return nil
}
