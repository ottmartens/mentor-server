package controllers

import (
	"encoding/json"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	//
	profile := models.AccountPublic{}

	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
	}

	user := models.GetUser(userId, false)

	if len(profile.FirstName) > 0 {
		user.FirstName = profile.FirstName
	}
	if len(profile.LastName) > 0 {
		user.LastName = profile.LastName
	}
	if len(profile.Bio) > 0 {
		user.Bio = profile.Bio
	}

	models.GetDB().Save(user)

	utils.Respond(w, utils.Message(true, "Profile successfully edited"))
}
