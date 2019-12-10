package test

import (
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/controllers"
	"github.com/ottmartens/mentor-server/models"
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupFormingRequest(t *testing.T) {

	testRouter := mux.NewRouter()
	testRouter.Use(models.JwtAuthentication)
	testRouter.HandleFunc("/api/groups/request-creation", controllers.RequestGroupForming).Methods("POST")

	server := httptest.NewServer(testRouter)
	defer server.Close()

	secondMentor := &models.Account{}
	models.GetDB().Table("accounts").Where("email = ?", TestEmail2).Find(secondMentor)

	e := httpexpect.New(t, server.URL)

	obj := e.POST("/api/groups/request-creation").
		WithHeader("Authorization", TestUserToken).
		WithJSON(map[string]interface{}{
			"userId": secondMentor.ID,
		}).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("success", "message")

	if obj.Value("success").Boolean().Raw() != true {
		t.Error("Requesting group forming failed")
	}
}
