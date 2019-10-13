package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
)

type Group struct {
	gorm.Model
	Title       string `json:"title"`
	Tagline     string `json:"tagline"`
	Description string `json:"description"`
}

type GroupWithMentors struct {
	gorm.Model
	Title       string   `json:"title"`
	Tagline     string   `json:"tagline"`
	Description string   `json:"description"`
	Mentors     []Mentor `json:"mentors"`
}

func GetGroups() []GroupWithMentors {

	groups := make([]*Group, 0)
	err := GetDB().Table("groups").Find(&groups).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var resp []GroupWithMentors

	for _, group := range groups {
		resp = append(resp, group.GetMentors())
	}

	return resp
}

func (group *Group) Create(mentors []uint) map[string]interface{} {

	if resp, ok := group.Validate(mentors); !ok {
		return resp
	}

	GetDB().Create(group)

	if group.ID <= 0 {
		return utils.Message(false, "Failed to create group, connection error")
	}

	for _, userId := range mentors {
		user := GetUser(userId)
		user.SetGroupId(group.ID)
	}

	response := utils.Message(true, "Group has been created")
	response["group"] = group
	return response
}

func (group *Group) Validate(mentors []uint) (map[string]interface{}, bool) {

	if len(mentors) == 0 || len(mentors) > 3 {
		return utils.Message(false, "The amount of mentors must be between 1-3"), false
	}

	for _, userId := range mentors {
		user := &Account{}
		err := GetDB().Table("accounts").Where("id = ?", userId).First(user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return utils.Message(false, "Invalid mentor ids"), false
			}
			return utils.Message(false, "Connection error, please try again"), false
		}

		if user.Role != enums.UserTypes.Mentor {
			return utils.Message(false, fmt.Sprintf("User %s is not a mentor", user.Email)), false
		}

		if user.GroupId > 0 {
			return utils.Message(false, fmt.Sprintf("User %s already belongs to a group", user.Email)), false
		}
	}

	return utils.Message(true, "Validation passed"), true
}

func (group *Group) GetMentors() GroupWithMentors {

	mentorAccounts := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("group_id = ?", group.ID).Find(&mentorAccounts).Error

	fmt.Println("len", len(mentorAccounts))

	if err != nil {
		fmt.Println(err)
	}

	response := GroupWithMentors{
		Title:       group.Title,
		Tagline:     group.Tagline,
		Description: group.Description,
		Mentors:     nil,
	}

	for _, account := range mentorAccounts {
		response.Mentors = append(response.Mentors, Mentor{
			FirstName: account.FirstName,
			LastName:  account.LastName,
			UserId:    account.ID,
			ImageUrl:  account.ImageUrl,
		})
	}

	return response
}
