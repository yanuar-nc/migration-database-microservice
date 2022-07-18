package usecase

import (
	"context"
	"fmt"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (u *UsecaseImplementation) Save(ctx context.Context, req domain.User) error {

	err := u.repository.Save(ctx, &req)
	if err != nil {
		return err
	}

	u.eventRepository.Publish(ctx, u.cfg.Event.TopicUser, domain.EventMessage{
		EventType: domain.EventInsert,
		Key:       fmt.Sprintf("%s", req.ID),
		Message:   req,
	})
	return nil
}

func (u *UsecaseImplementation) Update(ctx context.Context, req domain.User) error {

	err := u.repository.Update(ctx, &req)
	if err != nil {
		return err
	}

	u.eventRepository.Publish(ctx, u.cfg.Event.TopicUser, domain.EventMessage{
		EventType: domain.EventUpdate,
		Key:       fmt.Sprintf("%s", req.ID),
		Message:   req,
	})

	return nil
}
