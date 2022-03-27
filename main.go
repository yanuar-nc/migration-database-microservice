package main

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/go-boiler-plate/config"
	"github.com/yanuar-nc/go-boiler-plate/config/database"
	"github.com/yanuar-nc/go-boiler-plate/helper"
)

func main() {

	// call config.Load() before start up
	err := config.Load()
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "load_config")
		os.Exit(1)
	}

	writeDB, err := database.GetGormConn(config.WriteDBHost, config.WriteDBUser, config.WriteDBName, config.WriteDBPassword, config.WriteDBPort)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "write_db")
		os.Exit(1)
	}

	echoServer, err := NewEchoServer(writeDB, nil)
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
