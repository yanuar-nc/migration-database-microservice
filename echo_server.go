package main

import (
	"fmt"
	"net/http"

	"github.com/yanuar-nc/go-boiler-plate/config"

	deliveryJson "github.com/yanuar-nc/go-boiler-plate/src/delivery/json"
	gormRepository "github.com/yanuar-nc/go-boiler-plate/src/repository/gorm"
	usecasePackage "github.com/yanuar-nc/go-boiler-plate/src/usecase"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// EchoServer structure
type EchoServer struct {
	movieEchoHandler *deliveryJson.EchoHandler
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// movie v1 route
	movieGroupV1 := e.Group("/movie")
	s.movieEchoHandler.Mount(movieGroupV1)

	listenerPort := fmt.Sprintf(":%d", config.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(writeDb, readDb *gorm.DB) (*EchoServer, error) {
	repositoryImpl := gormRepository.NewRepository(writeDb)

	usecase := usecasePackage.NewUsecaseImplementation().PutRepository(repositoryImpl)

	movieEchoHandler := deliveryJson.NewEchoHandler(usecase)

	return &EchoServer{
		movieEchoHandler: movieEchoHandler,
	}, nil
}
