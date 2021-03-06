package gorm

import (
	"context"
	"errors"

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
	err := l.db.Debug().Save(data).Error
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

func (l *Repository) FetchAll(ctx context.Context, filter domain.Filter) ([]domain.User, error) {
	return nil, nil
}

func (l *Repository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var result *domain.User
	q := l.db.Debug().Where(&domain.User{ID: id}).First(&result)
	if q.Error != nil {
		if q.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(domain.DataIsNotFound)
		}
		return nil, q.Error
	}
	return &domain.User{}, nil
}
