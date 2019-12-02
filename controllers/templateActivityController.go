package controllers

import (
	"encoding/json"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func GetTemplateActivities(w http.ResponseWriter, _ *http.Request) {

	templateActivities := models.GetTemplateActivities()

	if templateActivities == nil {
		utils.Respond(w, utils.Message(false, "Could not get template activities"))
	}

	resp := utils.Message(true, "success")

	resp["data"] = templateActivities

	utils.Respond(w, resp)
}

func AddTemplateActivity(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	if !models.IsAdmin(userId) {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	templateActivity := models.TemplateActivity{}

	err := json.NewDecoder(r.Body).Decode(&templateActivity)

	if err != nil || len(templateActivity.Name) == 0 {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	templateActivity.Save()

	utils.Respond(w, utils.Message(true, "template activity saved!"))
}
