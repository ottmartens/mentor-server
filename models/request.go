package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
)

type Request struct {
	gorm.Model
	Type      string
	Initiator uint // Id of the requester
	Target    uint // Id of the target (group for joining, user for group creating, etc)
}

func CreateRequest(requestType string, initiatorId uint, targetId uint) map[string]interface{} {

	request := Request{
		Type:      requestType,
		Initiator: initiatorId,
		Target:    targetId,
	}

	if RequestExists(request) {
		return utils.Message(false, "This request already exists")
	}

	err := GetDB().Create(&request).Error
	if err != nil {
		return utils.Message(false, "Error creating request")
	}

	return utils.Message(true, "Request successfully created")
}

func RequestExists(request Request) bool {
	prev := &Request{}
	GetDB().Where(request).First(prev)
	return prev.ID != 0
}

func HandleJoiningRequest(groupId uint, userId uint, accept bool) map[string]interface{} {

	request := Request{
		Type:      enums.RequestTypes.JoinGroup,
		Initiator: userId,
		Target:    groupId,
	}

	GetDB().Where(&request).First(&request)

	if request.ID == 0 {
		return utils.Message(false, "Request not found")
	}

	GetDB().Delete(request)

	if accept {
		user := GetUser(userId, false)
		user.SetGroupId(groupId)
		return utils.Message(true, "Request approved!")
	}

	return utils.Message(true, "Request deleted!")
}

func HandleFormingRequest(userId uint, requesterId uint, accept bool) map[string]interface{} {

	request := Request{
		Type:      enums.RequestTypes.CreateGroup,
		Initiator: requesterId,
		Target:    userId,
	}

	GetDB().Where(&request).First(&request)

	if request.ID == 0 {
		return utils.Message(false, "Request not found")
	}

	if accept {
		mentorOne := GetUser(userId, false)
		mentorTwo := GetUser(requesterId, false)

		group := &Group{}

		if mentorOne.FirstName != "" && mentorTwo.LastName != "" {
			group.Tagline = fmt.Sprintf("The humble bundle of %s and %s", mentorOne.FirstName, mentorTwo.FirstName)
			group.Description = fmt.Sprintf("%s will get you drunk and %s is your key to passing your exams", mentorOne.FirstName, mentorTwo.FirstName)
		}

		err := GetDB().Create(group).Error
		if err != nil {
			return utils.Message(false, "Error creating the group")
		}

		mentorOne.SetGroupId(group.ID)
		mentorTwo.SetGroupId(group.ID)

		mentorOne.deleteAllFormingRequests()
		mentorTwo.deleteAllFormingRequests()

		resp := utils.Message(true, "Request approved!")
		resp["data"] = group
		return resp
	} else {
		GetDB().Delete(request)
		return utils.Message(true, "Request deleted!")
	}
}

func (account *Account) deleteAllFormingRequests() {

	requests := make([]*Request, 0)

	err := GetDB().Table("requests").Where(&Request{
		Type:      enums.RequestTypes.CreateGroup,
		Initiator: account.ID,
	}).Or(&Request{
		Type:   enums.RequestTypes.CreateGroup,
		Target: account.ID,
	}).Find(&requests).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("Error deleting requests for user %d \n", account.ID)
		fmt.Println(err)
	}

	for _, request := range requests {
		GetDB().Delete(request)
	}
	return
}
