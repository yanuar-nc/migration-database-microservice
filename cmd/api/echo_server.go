package main

import (
	"fmt"
	"net/http"

	"github.com/yanuar-nc/migration-database-microservice/config"

	deliveryJson "github.com/yanuar-nc/migration-database-microservice/src/delivery/json"
	firestoreRepository "github.com/yanuar-nc/migration-database-microservice/src/repository/firestore"
	kafkaRepository "github.com/yanuar-nc/migration-database-microservice/src/repository/kafka"
	usecasePackage "github.com/yanuar-nc/migration-database-microservice/src/usecase"

	"github.com/labstack/echo"
)

// EchoServer structure
type EchoServer struct {
	echoHandler *deliveryJson.EchoHandler
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// user v1 route
	userGroupV1 := e.Group("/user")
	s.echoHandler.Mount(userGroupV1)

	listenerPort := fmt.Sprintf(":%d", config.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(cfg config.Config) (*EchoServer, error) {
	firestoreRepository := firestoreRepository.NewRepository(cfg.Firestore.Client)
	kafkaRepository := kafkaRepository.NewKafka([]string{cfg.Event.Broker}, cfg.Event.Config)
	usecase := usecasePackage.NewUsecaseImplementation(cfg).
		PutRepository(firestoreRepository).
		PutEventRepository(kafkaRepository)

	echoHandler := deliveryJson.NewEchoHandler(usecase)

	return &EchoServer{
		echoHandler: echoHandler,
	}, nil
}
