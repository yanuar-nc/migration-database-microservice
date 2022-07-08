package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/go-boiler-plate/src/domain"
	"github.com/yanuar-nc/golang/helper"
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

func (l *Repository) Save(ctx context.Context, data *domain.Movie) error {
	data.Datetime = time.Now()
	err := l.db.Save(data).Error
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "Repository", "save")
		return err
	}
	return nil
}
