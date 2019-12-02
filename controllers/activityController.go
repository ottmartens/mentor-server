package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	. "github.com/ottmartens/mentor-server/utils/enums"
	"net/http"
)

func AddGroupActivity(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	user := models.GetUser(userId, false)

	if user.Role != UserTypes.Mentor || user.GroupId == nil {
		utils.Respond(w, utils.Message(false, "No groupID, or you are not a mentor"))
		return
	}

	type req struct {
		TemplateId   *uint    `json:"templateId"`
		Name         string   `json:"name"`
		Time         string   `json:"time"`
		Participants []int64  `json:"participants"`
		Images       []string `json:"images"`
	}

	payload := req{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil || len(payload.Time) == 0 || len(payload.Name) == 0 || len(payload.Images) == 0 {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	activity := models.Activity{}

	if payload.TemplateId != nil {

		templateActivity := models.GetTemplateActivity(*payload.TemplateId)

		if uint(len(payload.Participants)) < templateActivity.RequiredParticipants {
			utils.Respond(w, utils.Message(false, "This activity needs more participants"))
			return
		}

		activity.TemplateId = payload.TemplateId
		activity.Name = templateActivity.Name
		activity.Points = templateActivity.Points

	} else {
		activity.Name = payload.Name

		if len(payload.Participants) < 3 {
			utils.Respond(w, utils.Message(false, "Activity must have at least 3 participants"))
			return
		}
	}

	activity.GroupId = *user.GroupId
	activity.Participants = payload.Participants
	activity.Time = payload.Time
	activity.Images = payload.Images

	err = models.GetDB().Save(&activity).Error
	if err != nil {
		fmt.Println("Error adding activity:", err)
		utils.Respond(w, utils.Message(false, "Error adding activity"))
		return
	}

	utils.Respond(w, utils.Message(true, "Activity successfully added!"))
}
