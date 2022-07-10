package usecase

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (u *UsecaseImplementation) EventSave(ctx context.Context, req domain.User) error {

	err := u.repository.Save(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsecaseImplementation) EventUpdate(ctx context.Context, req domain.User) error {

	err := u.repository.Update(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
