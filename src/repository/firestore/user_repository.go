package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

// Repository struct
type Repository struct {
	client *firestore.Client
}

// NewRepository function
func NewRepository(client *firestore.Client) *Repository {
	return &Repository{client: client}
}

func (l *Repository) Save(ctx context.Context, data *domain.User) error {
	return nil
}

func (l *Repository) Update(ctx context.Context, data *domain.User) error {
	return nil
}

func (l *Repository) FetchAll(ctx context.Context, filter domain.Filter) ([]domain.User, error) {
	return nil, nil
}

func (l *Repository) MigrationUpdate(ctx context.Context, data *domain.Migration) error {
	return nil
}

func (l *Repository) MigrationGet(ctx context.Context) (*domain.Migration, error) {
	return &domain.Migration{}, nil
}

func (l *Repository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return &domain.User{}, nil
}
