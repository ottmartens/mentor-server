package controllers

import (
	"encoding/json"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	. "github.com/ottmartens/mentor-server/utils/enums"
	"net/http"
)

func isAdmin(r *http.Request) bool {
	userId := r.Context().Value("user").(uint)

	user := models.GetUser(userId, true)

	if user == nil || user.Role != UserTypes.Admin {
		return false
	} else {
		return true
	}
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {

	if !isAdmin(r) {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	type request struct {
		UserID       uint   `json:"userId"`
		Accept       bool   `json:"accept"`
		RejectReason string `json:"rejectReason"`
	}

	payload := request{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid payload"))
		return
	}

	user := models.GetUser(payload.UserID, false)

	user.IsVerified = &payload.Accept
	user.RejectionReason = payload.RejectReason

	models.GetDB().Save(user)

	utils.Respond(w, utils.Message(true, "success"))
}
