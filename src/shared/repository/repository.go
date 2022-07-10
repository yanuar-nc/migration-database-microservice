package repository

import (
	"context"

	"github.com/yanuar-nc/migration-database-microservice/src/shared/domain"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, event domain.EventMessage) error
}
