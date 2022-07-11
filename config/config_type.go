package config

import (
	"cloud.google.com/go/firestore"
	"github.com/Shopify/sarama"
	"gorm.io/gorm"
)

type ConfigWriteDB struct {
	// Host config
	Host string
	// Name config
	Name string
	// User config
	User string
	// Password config
	Password string
	// Port config
	Port int

	DB *gorm.DB
}

type ConfigReadDB struct {
	// Host config
	Host string
	// Name config
	Name string
	// User config
	User string
	// Password config
	Password string
	// Port config
	Port int
}

type EventConfig struct {
	// Broker config
	Broker string

	// GroupID config
	GroupID string

	// ClientID config
	ClientID string

	// Topics config
	Topics []string

	// TopicUser config
	TopicUser string

	Config *sarama.Config
}

func (c EventConfig) EventInsert() string {
	return "INSERT"
}

func (c EventConfig) EventUpdate() string {
	return "UPDATE"
}

type FirestoreConfig struct {
	ProjectID string
	Client    *firestore.Client
}
