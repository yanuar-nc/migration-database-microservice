package main

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/migration-database-microservice/config"
	"github.com/yanuar-nc/migration-database-microservice/config/database"
	"github.com/yanuar-nc/migration-database-microservice/config/messagebroker"
	"github.com/yanuar-nc/migration-database-microservice/helper"
)

func main() {

	// call config.Load() before start up
	cfg, err := config.Load()
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "load_config")
		os.Exit(1)
	}

	client, err := database.GetFirestoreConn(cfg.Firestore.ProjectID)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "client_db")
		os.Exit(1)
	}
	cfg.Firestore.Client = client

	kafkaClient := messagebroker.GetKafkaConfig(cfg.Event)
	cfg.Event.Config = kafkaClient

	echoServer, err := NewEchoServer(cfg)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "echo_server")
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		echoServer.Run()
	}()

	// Wait All services to end
	wg.Wait()
}
