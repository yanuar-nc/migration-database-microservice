package helper

import (
	"testing"

	"github.com/labstack/echo"
)

func TestNewRecorder(t *testing.T) {
	var payload = struct {
		name string
		res  string
	}{
		name: "test",
		res:  "testing",
	}
	_, _ = NewRecorder(t, "http://localhost/test", payload, echo.POST)
}
