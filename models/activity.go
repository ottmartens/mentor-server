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
