package test

import (
	"github.com/jinzhu/gorm"
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
	TestGroupID    uint
	db             *gorm.DB
)

func TestMain(m *testing.M) {
	db = models.GetDB() // Initialize DB connection
	m.Run()
	removeTestingData()
}

func removeTestingData() {
	testAccount := models.Account{}
	db.Table("accounts").Where("email = ?", TestEmail).Find(&testAccount)
	db.Unscoped().Delete(testAccount)
	db.Unscoped().Where("email = ?", TestEmail2).Delete(models.Account{})
	db.Unscoped().Where("initiator = ?", testAccount.ID).Delete(models.Request{})
	db.Unscoped().Where("id = ?", TestGroupID).Delete(models.Group{})

	TestUserToken = ""
	TestUserToken2 = ""
	TestGroupID = 0
}
