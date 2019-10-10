package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
)

type Group struct {
	gorm.Model
	Title       string `json:"title"`
	Tagline     string `json:"tagline"`
	Description string `json:"description"`
}

func GetGroups() []*Group {

	groups := make([]*Group, 0)
	err := GetDB().Table("groups").Find(&groups).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, group := range groups {
		fmt.Println(group.Title)
	}

	return groups
}

func (group *Group) Create() map[string]interface{} {

	GetDB().Create(group)

	if group.ID <= 0 {
		return utils.Message(false, "Failed to create group, connection error")
	}

	response := utils.Message(true, "Group has been created")
	response["group"] = group
	return response
}
