package test

import (
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/models"
	"testing"
)

func TestAccountPublicProfile(t *testing.T) {

	account := models.Account{
		Model: gorm.Model{
			ID: 1,
		},
		Email:     "Email",
		Password:  "PasswordHash",
		FirstName: "FirstName",
		LastName:  "LastName",
		Role:      "Role",
		Bio:       "Biography",
	}

	accountPublicProfile := models.AccountPublic{
		FirstName: "FirstName",
		LastName:  "LastName",
		UserId:    1,
		ImageUrl:  "",
		Bio:       "Biography",
	}

	result := account.GetPublicInfo()

	if result != accountPublicProfile {
		t.Error("account.GetPublicInfo returns invalid information")
	}
}
