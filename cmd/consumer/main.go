package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/migration-database-microservice/config"
	"github.com/yanuar-nc/migration-database-microservice/config/database"
	"github.com/yanuar-nc/migration-database-microservice/helper"
)

func main() {

	// call config.Load() before start up
	cfg, err := config.Load()
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "load_config")
		os.Exit(1)
	}

	writeDB, err := database.GetGormConn(cfg)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "write_db")
		os.Exit(1)
	}
	cfg.WriteDB.DB = writeDB

	err = NewEvent(cfg)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "echo_server")
		os.Exit(1)
	}

}
