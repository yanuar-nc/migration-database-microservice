package database

import (
	"fmt"

	"github.com/yanuar-nc/migration-database-microservice/config"
	domain "github.com/yanuar-nc/migration-database-microservice/src/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetGormConn function
func GetGormConn(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.WriteDB.Host, cfg.WriteDB.Port, cfg.WriteDB.User, cfg.WriteDB.Name, cfg.WriteDB.Password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&domain.User{}, &domain.Migration{})
	return db, err
}
