package main

import (
	"fmt"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/yanuar-nc/migration-database-microservice/config"
	"github.com/yanuar-nc/migration-database-microservice/helper"
	"github.com/yanuar-nc/migration-database-microservice/src/repository"

	"github.com/yanuar-nc/migration-database-microservice/src/delivery"
	"github.com/yanuar-nc/migration-database-microservice/src/usecase"
)

func NewEvent(writeDb, readDb *gorm.DB) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// cluster kafka construct with partitions mode
	clusterConfig := cluster.NewConfig()
	clusterConfig.ClientID = cfg.Kafka.ClientID
	clusterConfig.Group.Mode = cluster.ConsumerModePartitions

	// init consumer
	consumer, err := cluster.NewConsumer([]string{cfg.Kafka.Broker}, cfg.Kafka.GroupID, cfg.Kafka.Topics, clusterConfig)
	if err != nil {
		return err
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	movieRepositoryWrite := movieRepositoryPackage.NewMovieRepositoryGorm(writeDb)

	eventUsecase := eventUsecasePackage.NewEventUsecaseImpl(movieRepositoryWrite)
	eventEventHandler := eventDeliveryPackage.NewEventHandler(eventUsecase)

	for {
		select {
		case partition, ok := <-consumer.Partitions():
			if !ok {
				os.Exit(1)
			}

			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					switch msg.Topic {
					case cfg.Kafka.TopicMovie:
						err := eventEventHandler.Movie(msg.Key, msg.Value)
						if err != nil {
							helper.Log(log.ErrorLevel, err.Error(), "Event", "movie")
						}
						helper.Log(log.InfoLevel, "movie is successfully consumed", "Event", "movie")

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
