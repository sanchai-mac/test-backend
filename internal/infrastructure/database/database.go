package database

import (
	"log"
	"test-backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	CostomerDB *gorm.DB
	cfg        *config.Configuration
}

func NewDB(
	cfg *config.Configuration,
) *DB {
	db := DB{
		cfg: cfg,
	}

	// Connect to CustomerDB
	costomerDB, err := db.connectPostgres(cfg.CustomerDB, "CustomerDB")
	if err != nil {
		log.Fatalf("[DB:NewDB] ConnectPostgres (CustomerDB): %s\n", err)
	}
	db.CostomerDB = costomerDB
	log.Println("[DB:NewDB] Connected to CustomerDB successfully")

	return &db
}

// connectPostgres...
func (d *DB) connectPostgres(dsn, name string) (*gorm.DB, error) {
	log.Printf("[DB:connectPostgres] Connecting to %s: %s\n", name, dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[DB:connectPostgres] Ping %s error: %s\n", name, err)
		return nil, err
	}
	return db, nil
}
