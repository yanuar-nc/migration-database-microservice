package usecase

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/domain"
	"github.com/yanuar-nc/go-boiler-plate/src/repository"
)

// UsecaseImplementation struct
type UsecaseImplementation struct {
	repository repository.Repository
}

// NewUsecaseImplementation function
func NewUsecaseImplementation() *UsecaseImplementation {
	return &UsecaseImplementation{}
}

func (u *UsecaseImplementation) PutRepository(repo repository.Repository) *UsecaseImplementation {
	u.repository = repo
	return u
}

func (u *UsecaseImplementation) Save(ctx context.Context, req domain.Movie) error {

	err := u.repository.Save(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
