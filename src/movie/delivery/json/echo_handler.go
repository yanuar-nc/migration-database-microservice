package json

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yanuar-nc/go-boiler-plate/helper"
	"github.com/yanuar-nc/go-boiler-plate/src/movie/domain"
	"github.com/yanuar-nc/go-boiler-plate/src/movie/usecase"
)

// EchoHandler structure
type EchoHandler struct {
	movieUsecase usecase.MovieUsecase
}

// NewEchoHandler function
// Returns *EchoHandler
func NewEchoHandler(movieUsecase usecase.MovieUsecase) *EchoHandler {
	return &EchoHandler{movieUsecase: movieUsecase}
}

// Mount function
// Params : *echo.Group
func (h *EchoHandler) Mount(group *echo.Group) {
	group.POST("", h.Save)
}

// Save handler
func (h *EchoHandler) Save(c echo.Context) error {

	response := new(helper.JSONSchemaTemplate)

	param := domain.Movie{}

	err := c.Bind(&param)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Code = http.StatusBadRequest
		response.SetData(helper.Empty{})
		return response.ShowHTTPResponse(c)
	}

	err = h.movieUsecase.Save(c.Request().Context(), param)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Code = http.StatusBadRequest
		response.SetData(helper.Empty{})
		return response.ShowHTTPResponse(c)
	}

	response.Success = true
	response.Message = "Post Movie Response"
	response.Code = http.StatusOK

	return response.ShowHTTPResponse(c)
}
