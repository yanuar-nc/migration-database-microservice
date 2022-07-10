package usecase

import (
	"github.com/yanuar-nc/migration-database-microservice/config"
	"github.com/yanuar-nc/migration-database-microservice/src/repository"
)

// UsecaseImplementation struct
type UsecaseImplementation struct {
	repository          repository.Repository
	repositoryMigration repository.Repository
	eventRepository     repository.EventRepository
	cfg                 config.Config
}

// NewUsecaseImplementation function
func NewUsecaseImplementation(cfg config.Config) *UsecaseImplementation {
	return &UsecaseImplementation{
		cfg: cfg,
	}
}

func (u *UsecaseImplementation) PutRepository(repo repository.Repository) *UsecaseImplementation {
	u.repository = repo
	return u
}

func (u *UsecaseImplementation) PutRepositoryMigration(repo repository.Repository) *UsecaseImplementation {
	u.repositoryMigration = repo
	return u
}

func (u *UsecaseImplementation) PutEventRepository(repo repository.EventRepository) *UsecaseImplementation {
	u.eventRepository = repo
	return u
}
