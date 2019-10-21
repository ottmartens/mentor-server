package models

import (
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

	if accept {
		user := GetUser(userId, false)
		user.SetGroupId(groupId)
	}

	GetDB().Delete(request)

	return utils.Message(true, "Request approved!")
}
