package config

import (
	"errors"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

var (

	// Development env checking, this env for debug purpose
	Development string

	// HTTPPort config
	HTTPPort uint16

	// GRPCPort config
	GRPCPort string
)

type Config struct {
	// Development env checking, this env for debug purpose
	Development string

	// HTTPPort config
	HTTPPort uint16

	// GRPCPort config
	GRPCPort string

	WriteDB ConfigWriteDB

	ReadDB ConfigReadDB

	Event EventConfig

	Firestore FirestoreConfig
}

// Load function will load all config from environment variable
func Load() (Config, error) {

	var (
		c   Config
		err error
	)

	// load .env
	err = dotenv.Load(".env")
	if err != nil {
		return c, err
	}

	err = c.setDevelopmentMode()
	if err != nil {
		return c, err
	}

	err = c.setHTTPPort()
	if err != nil {
		return c, err
	}

	err = c.setReadDB()
	if err != nil {
		return c, err
	}

	err = c.setWriteDB()
	if err != nil {
		return c, err
	}

	err = c.setKafka()
	if err != nil {
		return c, err
	}

	err = c.setFirestore()
	if err != nil {
		return c, err
	}

	return c, err
}

func (c *Config) setDevelopmentMode() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	development, ok := os.LookupEnv("DEVELOPMENT")
	if !ok {
		return errors.New("DEVELOPMENT env is not loaded")
	}

	c.Development = development
	return nil
}

func (c *Config) setHTTPPort() error {
	httpPortStr, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		return errors.New("HTTP_PORT env is not loaded")
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return errors.New("HTTP_PORT env is not valid")
	}

	// set http port
	c.HTTPPort = uint16(httpPort)
	return nil
}

func (c *Config) setWriteDB() error {
	writeDBHost, ok := os.LookupEnv("WRITE_DB_HOST")
	if !ok {
		return errors.New("WRITE_DB_HOST env is not loaded")
	}

	// set WriteDBHost
	c.WriteDB.Host = writeDBHost

	writeDBName, ok := os.LookupEnv("WRITE_DB_NAME")
	if !ok {
		return errors.New("WRITE_DB_NAME env is not loaded")
	}

	// set WriteDBName
	c.WriteDB.Name = writeDBName

	writeDBUser, ok := os.LookupEnv("WRITE_DB_USER")
	if !ok {
		return errors.New("WRITE_DB_USER env is not loaded")
	}

	// set WriteDBUser
	c.WriteDB.User = writeDBUser

	writeDBPassword, ok := os.LookupEnv("WRITE_DB_PASSWORD")
	if !ok {
		return errors.New("WRITE_DB_PASSWORD env is not loaded")
	}

	// set WriteDBPassword
	c.WriteDB.Password = writeDBPassword

	writeDBPort, ok := os.LookupEnv("WRITE_DB_PORT")
	if !ok {
		return errors.New("WRITE_DB_PORT env is not loaded")
	}

	// set WriteDBPort
	c.WriteDB.Port, _ = strconv.Atoi(writeDBPort)

	return nil
}

func (c *Config) setReadDB() error {
	readDBHost, ok := os.LookupEnv("READ_DB_HOST")
	if !ok {
		return errors.New("READ_DB_HOST env is not loaded")
	}

	// set ReadDBHost
	c.ReadDB.Host = readDBHost

	readDBName, ok := os.LookupEnv("READ_DB_NAME")
	if !ok {
		return errors.New("READ_DB_NAME env is not loaded")
	}

	// set ReadDBName
	c.ReadDB.Name = readDBName

	readDBUser, ok := os.LookupEnv("READ_DB_USER")
	if !ok {
		return errors.New("READ_DB_USER env is not loaded")
	}

	// set ReadDBUser
	c.ReadDB.User = readDBUser

	readDBPassword, ok := os.LookupEnv("READ_DB_PASSWORD")
	if !ok {
		return errors.New("READ_DB_PASSWORD env is not loaded")
	}

	// set ReadDBPassword
	c.ReadDB.Password = readDBPassword

	readDBPort, ok := os.LookupEnv("READ_DB_PORT")
	if !ok {
		return errors.New("READ_DB_PORT env is not loaded")
	}

	// set ReadDBPort
	c.ReadDB.Port, _ = strconv.Atoi(readDBPort)

	return nil
}

func (c *Config) setKafka() error {
	// set KafkaBroker
	kafkaBroker, ok := os.LookupEnv("KAFKA_BROKER")
	if !ok {
		return errors.New("KAFKA_BROKER env is not loaded")
	}

	// set KafkaBroker
	c.Event.Broker = kafkaBroker

	// set KafkaGroupID
	kafkaGroupId, ok := os.LookupEnv("KAFKA_GROUP_ID")
	if !ok {
		return errors.New("KAFKA_GROUP_ID env is not loaded")
	}

	// set KafkaGroupID
	c.Event.GroupID = kafkaGroupId

	// set KafkaClientID
	kafkaClientID, ok := os.LookupEnv("KAFKA_CLIENT_ID")
	if !ok {
		return errors.New("KAFKA_CLIENT_ID env is not loaded")
	}
	c.Event.ClientID = kafkaClientID

	// set KafkaTopics
	kafkaTopicUser, ok := os.LookupEnv("KAFKA_TOPIC_USER")
	if !ok {
		return errors.New("KAFKA_TOPIC_USER env is not loaded")
	}
	c.Event.TopicUser = kafkaTopicUser

	// set KafkaTopics
	c.Event.Topics = []string{kafkaTopicUser}

	return nil
}

func (c *Config) setFirestore() error {
	// set ProjectID
	c.Firestore.ProjectID = "stockbit-api-dev"

	return nil
}
