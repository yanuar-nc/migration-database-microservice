package usecase

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

// Usecase interface
type Usecase interface {
	Save(ctx context.Context, param domain.User) error
	Update(ctx context.Context, param domain.User) error
}
