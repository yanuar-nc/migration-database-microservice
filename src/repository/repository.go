package repository

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

// Repository interface
type Repository interface {
	Save(ctx context.Context, data *domain.User) error
	Update(ctx context.Context, data *domain.User) error
	GetByID(ctx context.Context, id int) (domain.User, error)
	FetchAll(ctx context.Context, filter domain.Filter) ([]domain.User, error)
	MigrationUpdate(ctx context.Context, data *domain.Migration) error
	MigrationGet(ctx context.Context) (*domain.Migration, error)
}

type EventRepository interface {
	Publish(ctx context.Context, topic string, event domain.EventMessage) error
}
