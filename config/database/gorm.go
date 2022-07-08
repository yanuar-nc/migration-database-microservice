package database

import (
	"fmt"

	movie "github.com/yanuar-nc/go-boiler-plate/src/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetGormConn function
func GetGormConn(host, user, dbName, password string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&movie.Movie{})
	return db, err
}
