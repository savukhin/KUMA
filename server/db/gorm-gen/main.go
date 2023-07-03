package main

import (
	"server/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "host=localhost user=20624880 password=admin dbname=kuma port=5432 sslmode=disable TimeZone=Europe/London"
	postgresConn := postgres.Open(dsn)
	gormdb, err := gorm.Open(postgresConn)

	if err != nil {
		panic(err)
	}

	g.UseDB(gormdb) // reuse gorm db

	migrateModels := models.MigrateModels

	// Generate basic type-safe DAO API for structs
	g.ApplyBasic(
		migrateModels...,
	)

	gormdb.AutoMigrate(migrateModels...)

	// Generate the code
	g.Execute()
}
