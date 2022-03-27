package repository

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/movie/domain"
)

// Movie interface
type Movie interface {
	Save(ctx context.Context, data *domain.Movie) error
}
