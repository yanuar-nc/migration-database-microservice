package usecase

import (
	"context"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/labstack/gommon/log"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (u *UsecaseImplementation) Migration(ctx context.Context) error {
	var (
		createdAt int
		err       error
	)

	m, err := u.repositoryMigration.MigrationGet(ctx)
	if err != nil {
		return err
	}

	users, err := u.migrationFetchAll(ctx, m.Version)
	if err != nil {
		return err
	}

	for _, user := range users {

		createdAt = int(user.CreatedAt.Unix())
		us, err := u.repositoryMigration.GetByID(ctx, user.ID)
		if err != nil && err.Error() == domain.DataIsNotFound {
			err = retry.Do(
				func() error {
					if err := u.repositoryMigration.Save(ctx, &user); err != nil {
						return err
					}
					return nil
				},
				retry.Attempts(3),
				retry.Delay(time.Second*1),
				retry.LastErrorOnly(true),
			)
			if err != nil {
				return err
			}
		}
		_ = us
		u.migrationUpdate(ctx, createdAt, nil)

	}
	return nil
}

func (u *UsecaseImplementation) migrationFetchAll(ctx context.Context, createdAt int64) ([]domain.User, error) {
	users, err := u.repository.FetchAll(ctx, domain.Filter{
		CreatedAt: domain.FilterValue{
			Sort:   "ASC",
			Value:  createdAt,
			Bigger: true,
		},
		Limit: 10,
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UsecaseImplementation) migrationUpdate(ctx context.Context, createdAt int, err error) {
	if err == nil {
		err = u.repositoryMigration.MigrationUpdate(ctx, &domain.Migration{
			Version:   int64(createdAt),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			log.Error(err)
		}
	}

}
