package models

import "github.com/jinzhu/gorm"

type Activity struct {
	gorm.Model
	GroupId         uint     `json:"groupId"`
	TemplateId      *uint    `json:"templateId"`
	Name            string   `json:"name"`
	Points          uint     `json:"points"`
	Time            string   `json:"time"`
	Participants    []uint   `json:"participants"`
	Images          []string `json:"images"`
	IsVerified      *bool    `json:"isVerified"`
	RejectionReason string   `json:"rejectionReason"`
}
