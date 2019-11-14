package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils/enums"
)

type AccountPublic struct {
	Name     string `json:"name"`
	Tagline  string `json:"tagline"`
	Degree   string `json:"degree"`
	Year     string `json:"year"`
	UserId   uint   `json:"userId"`
	ImageUrl string `json:"imageUrl"`
	Bio      string `json:"bio"`
}

type AvailableMentor struct {
	Name             string `json:"name"`
	UserId           uint   `json:"userId"`
	ImageUrl         string `json:"imageUrl"`
	HasRequestedYou  bool   `json:"hasRequestedYou"`
	YouHaveRequested bool   `json:"youHaveRequested"`
}

func GetFreeMentors(userId uint) []AvailableMentor {

	freeMentors := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("role = ?", enums.UserTypes.Mentor).Where("group_id IS NULL").Not("id = ?", userId).Find(&freeMentors).Error

	if err != nil {
		return nil
	}

	availableMentors := make([]AvailableMentor, 0)

	for _, mentor := range freeMentors {
		availableMentors = append(availableMentors, mentor.getGroupRequests(userId))
	}
	return availableMentors
}

func (account *Account) getGroupRequests(userId uint) AvailableMentor {
	yourRequest := Request{}
	requestToYou := Request{}

	createGroupRequestsQuery := GetDB().Table("requests").Where("type = ?", enums.RequestTypes.CreateGroup)

	err := createGroupRequestsQuery.Where("initiator = ?", userId).Where("target = ?", account.ID).First(&yourRequest).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}
	err = createGroupRequestsQuery.Where("initiator = ?", account.ID).Where("target = ?", userId).First(&requestToYou).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}

	return AvailableMentor{
		Name:             account.Name,
		UserId:           account.ID,
		ImageUrl:         account.ImageUrl,
		YouHaveRequested: yourRequest.ID != 0,
		HasRequestedYou:  requestToYou.ID != 0,
	}
}
