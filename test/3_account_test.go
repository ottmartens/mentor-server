package test

import (
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/controllers"
	"github.com/ottmartens/mentor-server/models"
	. "github.com/ottmartens/mentor-server/utils/enums"
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountPublicProfile(t *testing.T) {

	account := models.Account{
		Model: gorm.Model{
			ID: 1,
		},
		Email:    "Email",
		Password: "PasswordHash",
		Name:     "FirstName LastName",
		Role:     "Role",
		Bio:      "Biography",
	}

	accountPublicProfile := models.AccountPublic{
		Name:     "FirstName LastName",
		UserId:   1,
		ImageUrl: "",
		Bio:      "Biography",
	}

	result := account.GetPublicInfo()

	if result != accountPublicProfile {
		t.Error("account.GetPublicInfo returns invalid information")
	}
}

func TestAccountCreation(t *testing.T) {

	handler := http.HandlerFunc(controllers.CreateAccount)

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	obj := e.POST("").WithJSON(map[string]interface{}{
		"email":    TestEmail,
		"password": TestPassword,
		"role":     UserTypes.Mentor,
	}).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("success", "data", "message")
	obj.Value("data").Object().Keys().ContainsOnly("name", "imageUrl", "token", "role")

	TestUserToken = obj.Value("data").Object().Value("token").String().Raw()

	// 2nd account
	obj2 := e.POST("").WithJSON(map[string]interface{}{
		"email":    TestEmail2,
		"password": TestPassword,
		"role":     UserTypes.Mentor,
	}).Expect().
		Status(http.StatusOK).JSON().Object()

	TestUserToken2 = obj2.Value("data").Object().Value("token").String().Raw()
}
