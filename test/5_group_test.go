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

func TestGroupCreation(t *testing.T) {

	testRouter := mux.NewRouter()
	testRouter.Use(models.JwtAuthentication)
	testRouter.HandleFunc("/api/groups/accept-creation", controllers.HandleForming).Methods("POST")

	server := httptest.NewServer(testRouter)
	defer server.Close()

	firstMentor := &models.Account{}
	models.GetDB().Table("accounts").Where("email = ?", TestEmail).First(firstMentor)

	e := httpexpect.New(t, server.URL)

	obj := e.POST("/api/groups/accept-creation").
		WithHeader("Authorization", TestUserToken2).
		WithJSON(map[string]interface{}{
			"userId": firstMentor.ID,
			"accept": true,
		}).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("success", "message")

	models.GetDB().Table("accounts").Where("email = ?", TestEmail).First(firstMentor)

	if firstMentor.GroupId == nil {
		t.Error("GroupId not present on user")
	}

	TestGroupID = *firstMentor.GroupId
}
