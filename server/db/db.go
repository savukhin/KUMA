package db

import (
	"fmt"
	"server/config"
	"server/db/models"
	"server/db/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func autoMigrate(gormdb *gorm.DB) error {
	return gormdb.AutoMigrate(models.MigrateModels...)
}

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser,
		cfg.PostgresPassword, cfg.PostgresDB,
		cfg.PostgresPort,
	)

	postgresConn := postgres.Open(dsn)
	gormDB, err := gorm.Open(postgresConn)

	if err != nil {
		return gormDB, err
	}

	if cfg.PostgresAutoMigrate {
		err := autoMigrate(gormDB)
		if err != nil {
			return nil, err
		}
	}

	return gormDB, nil
}

func InitCncStatusTable(db *gorm.DB) {
	cnc := query.Use(db).CncStatus
	cnc.Create(models.StoppedStatus)
	cnc.Create(models.WorkingStatus)
	cnc.Create(models.BrokenStatus)
}
