package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Activity struct {
	gorm.Model
	GroupId         uint           `json:"groupId"`
	TemplateId      *uint          `json:"templateId"`
	Name            string         `json:"name"`
	Points          uint           `json:"points"`
	Time            string         `json:"time"`
	Participants    pq.Int64Array  `json:"participants" gorm:"type:int[]"`
	Images          pq.StringArray `json:"images" gorm:"type:varchar(100)[]"`
	IsVerified      *bool          `json:"isVerified"`
	RejectionReason string         `json:"rejectionReason"`
}

type UnverifiedActivity struct {
	Name      string `json:"name"`
	ID        uint
	GroupName string `json:"groupName"`
}

func GetActivity(id uint) *Activity {
	activity := &Activity{}

	GetDB().Table("activities").Where("id = ?", id).First(activity)
	return activity
}

func GetUnverifiedActivities() []UnverifiedActivity {
	activities := make([]Activity, 0)

	GetDB().Table("activities").Where("is_verified IS NULL").Find(&activities)

	unverifiedActivities := make([]UnverifiedActivity, 0)

	for _, activity := range activities {
		group := GetGroup(activity.GroupId)

		unverifiedActivities = append(unverifiedActivities, UnverifiedActivity{
			Name:      activity.Name,
			ID:        activity.ID,
			GroupName: group.Title,
		})
	}

	return unverifiedActivities
}
