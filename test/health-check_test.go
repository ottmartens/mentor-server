package test

import (
	"github.com/ottmartens/mentor-server/utils"
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	handler := http.HandlerFunc(utils.HealthCheck)

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	obj := e.GET("/api/health").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("success", "message")

	obj.Value("success").Boolean()
	obj.Value("message").String()
}
