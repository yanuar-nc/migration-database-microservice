package usecase

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/domain"
)

// Usecase interface
type Usecase interface {
	Save(ctx context.Context, param domain.Movie) error
}
