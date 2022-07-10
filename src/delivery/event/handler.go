package delivery

import (
	"context"
	"encoding/json"

	"github.com/yanuar-nc/migration-database-microservice/src/domain"
	"github.com/yanuar-nc/migration-database-microservice/src/usecase"
)

// EventHandler structure
type EventHandler struct {
	usecase usecase.Usecase
}

// NewEventHandler function
// Returns *EventHandler
func NewEventHandler(usecase usecase.Usecase) *EventHandler {
	return &EventHandler{usecase: usecase}
}

func (c *EventHandler) User(key, message []byte) error {
	var param domain.EventMessage
	if err := json.Unmarshal(message, &param); err != nil {
		return err
	}
	byteData, _ := json.Marshal(param.Message)
	switch param.EventType {
	case domain.EventInsert:
		err := c.Create(key, byteData)
		if err != nil {
			return err
		}
		return nil
	case domain.EventUpdate:
		err := c.Update(key, byteData)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (c *EventHandler) Create(key, message []byte) error {

	var param domain.User
	if err := json.Unmarshal(message, &param); err != nil {
		return err
	}

	err := c.usecase.Save(context.Background(), param)
	if err != nil {
		return err
	}
	return nil

}

func (c *EventHandler) Update(key, message []byte) error {

	var param domain.User
	if err := json.Unmarshal(message, &param); err != nil {
		return err
	}

	err := c.usecase.Update(context.Background(), param)
	if err != nil {
		return err
	}
	return nil

}
