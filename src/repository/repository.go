package repository

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/domain"
)

// Repository interface
type Repository interface {
	Save(ctx context.Context, data *domain.Movie) error
}
