package main

import (
	"fmt"
	"net/http"

	"github.com/yanuar-nc/go-boiler-plate/config"

	movieDeliveryJson "github.com/yanuar-nc/go-boiler-plate/src/movie/delivery/json"
	movieRepository "github.com/yanuar-nc/go-boiler-plate/src/movie/repository"
	movieUsecasePackage "github.com/yanuar-nc/go-boiler-plate/src/movie/usecase"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// EchoServer structure
type EchoServer struct {
	movieEchoHandler *movieDeliveryJson.EchoHandler
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
	movieRepositoryImpl := movieRepository.NewMovieRepositoryGorm(writeDb)

	movieUsecase := movieUsecasePackage.NewMovieUsecaseImpl(movieRepositoryImpl)

	movieEchoHandler := movieDeliveryJson.NewEchoHandler(movieUsecase)

	return &EchoServer{
		movieEchoHandler: movieEchoHandler,
	}, nil
}
