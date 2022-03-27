package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/go-boiler-plate/helper"
	"github.com/yanuar-nc/go-boiler-plate/src/movie/domain"
	"gorm.io/gorm"
)

// MovieRepositoryGorm struct
type MovieRepositoryGorm struct {
	db *gorm.DB
}

// NewMovieRepositoryGorm function
func NewMovieRepositoryGorm(db *gorm.DB) *MovieRepositoryGorm {
	return &MovieRepositoryGorm{db: db}
}

func (l *MovieRepositoryGorm) Save(ctx context.Context, data *domain.Movie) error {
	data.Datetime = time.Now()
	err := l.db.Save(data).Error
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "MovieRepository", "save")
		return err
	}
	return nil
}
