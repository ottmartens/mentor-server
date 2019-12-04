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
		utils.Respond(w, utils.Message(false, "No groupId, or you are not a mentor"))
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

func VerifyActivity(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)
	if !models.IsAdmin(userId) {
		utils.Respond(w, utils.Message(false, "Not permitted"))
		return
	}

	type request struct {
		ActivityId      uint   `json:"id"`
		Accept          bool   `json:"accept"`
		RejectionReason string `json:"rejectionReason"`
		Points          int    `json:"points"`
	}

	payload := request{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	activity := models.GetActivity(payload.ActivityId)

	if activity == nil {
		utils.Respond(w, utils.Message(false, "Activity not found"))
		return
	}

	if payload.Accept {

	} else {

	}
}

func GetUnverifiedActivities(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	if !models.IsAdmin(userId) {
		utils.Respond(w, utils.Message(false, "Not permitted"))
		return
	}

	activities := models.GetUnverifiedActivities()

	resp := utils.Message(true, "success")

	resp["data"] = activities
	utils.Respond(w, resp)
}
