package models

import "github.com/jinzhu/gorm"

type Activity struct {
	gorm.Model
	GroupId         string   `json:"groupId"`
	TemplateId      string   `json:"templateId"`
	Name            string   `json:"name"`
	Points          uint     `json:"points"`
	Time            string   `json:"time"`
	Participants    []uint   `json:"participants"`
	Images          []string `json:"images"`
	IsVerified      *bool    `json:"isVerified"`
	RejectionReason string   `json:"rejectionReason"`
}
