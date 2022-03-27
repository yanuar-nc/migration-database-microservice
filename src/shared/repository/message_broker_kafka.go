package repository

import (
	"gorm.io/gorm"
)

// KafkaMessageBroker struct
type KafkaMessageBroker struct {
	db *gorm.DB
}

// NewKafkaMessageBroker function
func NewKafkaMessageBroker(db *gorm.DB) *KafkaMessageBroker {
	return &KafkaMessageBroker{db: db}
}

func (l *KafkaMessageBroker) Publish(req interface{}) error {
	return nil
}
