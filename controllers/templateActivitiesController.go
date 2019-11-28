package controllers

import (
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func GetTemplateActivities(w http.ResponseWriter, r *http.Request) {

	templateActivities := models.GetTemplateActivities()

	if templateActivities == nil {
		utils.Respond(w, utils.Message(false, "Could not get template activities"))
	}

	resp := utils.Message(true, "success")

	resp["data"] = templateActivities

	utils.Respond(w, resp)
}
