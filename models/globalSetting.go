package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
	. "github.com/ottmartens/mentor-server/utils/enums"
)

type GlobalSetting struct {
	gorm.Model
	Key       string
	Value     int
	IsBoolean bool
}

func (globalSetting *GlobalSetting) GetValue() interface{} {
	if globalSetting.IsBoolean {
		if globalSetting.Value == 0 {
			return false
		} else {
			return true
		}
	}

	return globalSetting.Value
}

func GetGlobalSettings() map[string]interface{} {
	resp := map[string]interface{}{}

	globalSettings := make([]GlobalSetting, 0)

	GetDB().Table("global_settings").Find(&globalSettings)

	for _, setting := range globalSettings {
		resp[setting.Key] = setting.GetValue()
	}

	return resp
}

func SaveGlobalSettings(settings map[string]interface{}) map[string]interface{} {

	for setting, value := range settings {

		globalSetting := &GlobalSetting{}

		err := GetDB().Table("global_settings").Where("key = ?", setting).Find(&globalSetting).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			return utils.Message(false, "error setting the company settings")
		}

		switch value.(type) {
		case bool:
			if globalSetting.ID > 0 && !globalSetting.IsBoolean {
				return utils.Message(false, "Cannot set a boolean value to integer field")

			}
			globalSetting.IsBoolean = true
			if value.(bool) {
				globalSetting.Value = 1
			} else {
				globalSetting.Value = 0
			}
		case float64:
			if globalSetting.ID > 0 && globalSetting.IsBoolean {
				return utils.Message(false, "Cannot set an integer value to boolean field")
			}
			globalSetting.IsBoolean = false
			globalSetting.Value = int(value.(float64))
		default:
			return utils.Message(false, "Value of "+setting+" must either be boolean or integer")
		}

		globalSetting.Key = setting

		GetDB().Save(&globalSetting)
	}

	return utils.Message(true, "Global settings successfully saved!")
}

func InitializeGlobalSettings() {
	fmt.Print("Initialising global settings.. ")

	mentorsCanRegister := &GlobalSetting{}

	err := GetDB().Table("global_settings").Where("key = ?", GlobalSettingsTypes.MentorsCanRegister).Find(mentorsCanRegister).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		mentorsCanRegister.Key = GlobalSettingsTypes.MentorsCanRegister
		mentorsCanRegister.IsBoolean = true
		mentorsCanRegister.Value = 1

		GetDB().Save(mentorsCanRegister)
	}

	menteesCanRegister := &GlobalSetting{}

	GetDB().Table("global_settings").Where("key = ?", GlobalSettingsTypes.MenteesCanRegister).Find(menteesCanRegister)
	if err != nil && err == gorm.ErrRecordNotFound {
		menteesCanRegister.Key = GlobalSettingsTypes.MenteesCanRegister
		menteesCanRegister.IsBoolean = true
		menteesCanRegister.Value = 1

		GetDB().Save(menteesCanRegister)
	}

	fmt.Println("done!")
}
