package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/models"
	"testing"
)

func TestAccountPublcProfile(t *testing.T) {

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
		FirstName: "Email",
		LastName:  "PasswordHash",
		UserId:    1,
		ImageUrl:  "",
		Bio:       "Biography",
	}

	result := account.GetPublicInfo()

	if result != accountPublicProfile {
		t.Error("account.GetPublicInfo returns invalid information")
	}
}
