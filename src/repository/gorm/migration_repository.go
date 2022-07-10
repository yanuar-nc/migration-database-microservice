package gorm

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (r *Repository) MigrationUpdate(ctx context.Context, data *domain.Migration) error {
	return nil
}

func (r *Repository) MigrationGet(ctx context.Context) (*domain.Migration, error) {
	return nil, nil
}
