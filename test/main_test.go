package test

import (
	"github.com/ottmartens/mentor-server/models"
	"testing"
)

const (
	TestEmail    = "testaccount@mail.com"
	TestEmail2   = "testaccount2@mail.com"
	TestPassword = "password"
)

var (
	TestUserToken  string
	TestUserToken2 string
)

func TestMain(m *testing.M) {
	_ = models.GetDB() // Initialize DB connection
	m.Run()
	removeTestingData()
}

func removeTestingData() {
	testAccount := models.Account{}
	models.GetDB().Table("accounts").Where("email = ?", TestEmail).Find(&testAccount)
	models.GetDB().Unscoped().Delete(testAccount)
	models.GetDB().Unscoped().Where("email = ?", TestEmail2).Delete(models.Account{})
	models.GetDB().Unscoped().Where("initiator = ?", testAccount.ID).Delete(models.Request{})
	TestUserToken = ""
	TestUserToken2 = ""
}
