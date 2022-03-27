package usecase

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/movie/domain"
	"github.com/yanuar-nc/go-boiler-plate/src/movie/repository"
)

// MovieUsecaseImpl struct
type MovieUsecaseImpl struct {
	repository repository.Movie
}

// NewMovieUsecaseImpl function
func NewMovieUsecaseImpl(movieRepository repository.Movie) *MovieUsecaseImpl {
	return &MovieUsecaseImpl{repository: movieRepository}
}

func (u *MovieUsecaseImpl) Save(ctx context.Context, req domain.Movie) error {

	err := u.repository.Save(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
