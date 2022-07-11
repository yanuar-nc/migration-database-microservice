package main

import (
	"fmt"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
	log "github.com/sirupsen/logrus"
	delivery "github.com/yanuar-nc/migration-database-microservice/src/delivery/event"
	repository "github.com/yanuar-nc/migration-database-microservice/src/repository/gorm"
	kafkaRepository "github.com/yanuar-nc/migration-database-microservice/src/repository/kafka"
	usecasePackage "github.com/yanuar-nc/migration-database-microservice/src/usecase"

	"github.com/yanuar-nc/migration-database-microservice/config"
	"github.com/yanuar-nc/migration-database-microservice/helper"
)

func NewEvent(cfg config.Config) error {

	// cluster kafka construct with partitions mode
	clusterConfig := cluster.NewConfig()
	clusterConfig.ClientID = cfg.Event.ClientID
	clusterConfig.Group.Mode = cluster.ConsumerModePartitions

	// init consumer
	consumer, err := cluster.NewConsumer([]string{cfg.Event.Broker}, cfg.Event.GroupID, cfg.Event.Topics, clusterConfig)
	if err != nil {
		return err
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	firestoreRepository := repository.NewRepository(cfg.WriteDB.DB)
	kafkaRepository := kafkaRepository.NewKafka([]string{cfg.Event.Broker}, cfg.Event.Config)
	usecase := usecasePackage.NewUsecaseImplementation(cfg).
		PutRepository(firestoreRepository).
		PutEventRepository(kafkaRepository)

	eventHandler := delivery.NewEventHandler(usecase)

	for {
		select {
		case partition, ok := <-consumer.Partitions():
			if !ok {
				os.Exit(1)
			}

			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					switch msg.Topic {
					case cfg.Event.TopicUser:
						err := eventHandler.User(msg.Key, msg.Value)
						if err != nil {
							helper.Log(log.ErrorLevel, err.Error(), "Event", "user")
						}
						helper.Log(log.InfoLevel, "user is successfully consumed", "Event", "user")

					}

					//mark message as processed
					consumer.MarkOffset(msg, fmt.Sprintf("Key %s and offset %d is successfully consumed", msg.Key, msg.Offset))

				}
			}(partition)

		case <-signals:
			os.Exit(1)
		}
	}
}
