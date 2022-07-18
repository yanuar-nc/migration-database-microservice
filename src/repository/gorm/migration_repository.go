package gorm

import (
	"context"
	"time"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
	"gorm.io/gorm"
)

func (r *Repository) MigrationUpdate(ctx context.Context, data *domain.Migration) error {
	update := r.db.Debug().Model(domain.Migration{ID: 1}).Updates(data)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (r *Repository) MigrationGet(ctx context.Context) (*domain.Migration, error) {

	var data domain.Migration
	get := r.db.First(&data)
	if get.Error != nil {
		if get.Error == gorm.ErrRecordNotFound {
			data = domain.Migration{
				Version:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			create := r.db.Create(&data)
			if create.Error != nil {
				return nil, create.Error
			}
			return &data, nil
		}
		return nil, get.Error
	}

	return &data, nil
}
