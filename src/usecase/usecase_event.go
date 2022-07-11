package usecase

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (u *UsecaseImplementation) EventSave(ctx context.Context, req domain.User) error {

	user, err := u.repository.GetByID(ctx, req.ID)
	if err != nil && err.Error() != domain.DataIsNotFound {
		return err
	}

	if user == nil {
		err = u.repository.Save(ctx, &req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UsecaseImplementation) EventUpdate(ctx context.Context, req domain.User) error {

	err := u.EventSave(ctx, req)
	if err != nil {
		return err
	}

	err = u.repository.Update(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
