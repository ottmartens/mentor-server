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
	Id          uint            `json:"id"`
	Title       string          `json:"title"`
	Tagline     string          `json:"tagline"`
	Description string          `json:"description"`
	Mentors     []AccountPublic `json:"mentors"`
}

type GroupDetails struct {
	Title       string          `json:"title"`
	Tagline     string          `json:"tagline"`
	Description string          `json:"description"`
	Mentors     []AccountPublic `json:"mentors"`
	Mentees     []AccountPublic `json:"mentees"`
	Requests    []AccountPublic `json:"requests"`
}

func GetGroup(id uint) *Group {

	group := &Group{}
	GetDB().Table("groups").Where("id = ?", id).First(group)
	if group.ID == 0 {
		fmt.Printf("No group with id %d", id)
		return nil
	}

	return group
}

func GetGroupDetails(groupId uint) *GroupDetails {

	group := GetGroup(groupId)
	if group == nil {
		return nil
	}

	groupDetails := group.GetDetails()

	return &groupDetails
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
		user := GetUser(userId, false)
		user.SetGroupId(group.ID)
	}

	response := utils.Message(true, "Group has been created")
	response["data"] = group
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

		if *user.GroupId > 0 {
			return utils.Message(false, fmt.Sprintf("User %s already belongs to a group", user.Email)), false
		}
	}

	return utils.Message(true, "Validation passed"), true
}

func (group *Group) GetMentors() GroupWithMentors {

	mentors := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("group_id = ?", group.ID).Where("role = ?", enums.UserTypes.Mentor).Find(&mentors).Error

	if err != nil {
		fmt.Println(err)
	}

	response := GroupWithMentors{
		Title:       group.Title,
		Tagline:     group.Tagline,
		Description: group.Description,
		Id:          group.ID,
		Mentors:     nil,
	}

	for _, mentor := range mentors {
		response.Mentors = append(response.Mentors, mentor.getPublicInfo())
	}

	return response
}

func (group *Group) GetDetails() GroupDetails {

	groupDetails := GroupDetails{
		Title:       group.Title,
		Tagline:     group.Tagline,
		Description: group.Description,
		Mentors:     nil,
		Mentees:     nil,
		Requests:    nil,
	}

	mentors := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("group_id = ?", group.ID).Where("role = ?", enums.UserTypes.Mentor).Find(&mentors).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, mentor := range mentors {
		groupDetails.Mentors = append(groupDetails.Mentors, mentor.getPublicInfo())
	}

	mentees := make([]*Account, 0)
	err = GetDB().Table("accounts").Where("group_id = ?", group.ID).Where("role = ?", enums.UserTypes.Mentee).Find(&mentees).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, mentee := range mentees {
		groupDetails.Mentees = append(groupDetails.Mentees, mentee.getPublicInfo())
	}

	joiningRequests := make([]*Request, 0)
	err = GetDB().Table("requests").Where("target = ?", group.ID).Where("type = ?", enums.RequestTypes.JoinGroup).Find(&joiningRequests).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, request := range joiningRequests {
		user := GetUser(request.Initiator, true)
		groupDetails.Requests = append(groupDetails.Requests, user.getPublicInfo())
	}

	return groupDetails
}
