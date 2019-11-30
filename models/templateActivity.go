package models

import "github.com/jinzhu/gorm"

type TemplateActivity struct {
	gorm.Model
	Name                 string `json:"name"`
	Points               uint   `json:"points"`
	RequiredParticipants uint   `json:"requiredParticipants"`
}

func GetTemplateActivities() []TemplateActivity {
	templateActivities := make([]TemplateActivity, 0)
	_ = GetDB().Table("template_activities").Find(&templateActivities).Error

	return templateActivities
}

func (templateActivity *TemplateActivity) Save() {
	GetDB().Save(templateActivity)
}
