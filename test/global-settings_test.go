package test

import (
	"github.com/ottmartens/mentor-server/controllers"
	. "github.com/ottmartens/mentor-server/utils/enums"
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGlobalSettingsHandler(t *testing.T) {

	handler := http.HandlerFunc(controllers.GetGlobalSettings)

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	obj := e.GET("/api/global-settings").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("success", "message", "data")
	obj.Value("success").Boolean()
	obj.Value("message").String()

	data := obj.Value("data").Object()
	data.ContainsKey(GlobalSettingsTypes.MenteesCanRegister).ContainsKey(GlobalSettingsTypes.MentorsCanRegister)

	data.Value(GlobalSettingsTypes.MentorsCanRegister).Boolean()
	data.Value(GlobalSettingsTypes.MenteesCanRegister).Boolean()
}
