package usecase

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

func (u *UsecaseImplementation) Save(ctx context.Context, req domain.User) error {

	u4, _ := uuid.NewV4()
	req.ID = u4.String()
	req.CreatedAt = time.Now()
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
