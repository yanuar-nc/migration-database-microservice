package gorm

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/golang/helper"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	db *gorm.DB
}

// NewRepository function
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (l *Repository) Save(ctx context.Context, data *domain.User) error {
	err := l.db.Save(data).Error
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "Repository", "save")
		return err
	}
	return nil
}

func (l *Repository) Update(ctx context.Context, data *domain.User) error {
	err := l.db.Save(data).Error
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "Repository", "update")
		return err
	}
	return nil
}

func (l *Repository) GetByID(ctx context.Context, id int) (domain.User, error) {
	return domain.User{}, nil
}
