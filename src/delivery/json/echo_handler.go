package json

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yanuar-nc/migration-database-microservice/helper"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
	"github.com/yanuar-nc/migration-database-microservice/src/usecase"
)

// EchoHandler structure
type EchoHandler struct {
	usecase usecase.Usecase
}

// NewEchoHandler function
// Returns *EchoHandler
func NewEchoHandler(usecase usecase.Usecase) *EchoHandler {
	return &EchoHandler{usecase: usecase}
}

// Mount function
// Params : *echo.Group
func (h *EchoHandler) Mount(group *echo.Group) {
	group.POST("", h.Save)
}

// Save handler
func (h *EchoHandler) Save(c echo.Context) error {

	response := new(helper.JSONSchemaTemplate)

	param := domain.User{}

	err := c.Bind(&param)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Code = http.StatusBadRequest
		response.SetData(helper.Empty{})
		return response.ShowHTTPResponse(c)
	}

	err = h.usecase.Save(c.Request().Context(), param)
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
