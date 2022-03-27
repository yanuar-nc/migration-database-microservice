package helper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func NewRecorder(t *testing.T, url string, payload interface{}, method string) (echo.Context, *httptest.ResponseRecorder) {
	bodyData, err := json.Marshal(payload)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(method, url, strings.NewReader(string(bodyData)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
