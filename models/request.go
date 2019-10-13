package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
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
