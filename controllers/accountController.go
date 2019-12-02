package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	userIdInt, err := strconv.Atoi(mux.Vars(r)["id"])
	requesterId := r.Context().Value("user").(uint)
	userId := uint(userIdInt)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}

	if !models.IsAdmin(userId) && userId != requesterId {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	user := models.GetUser(uint(userId), false)

	if user == nil {
		utils.Respond(w, utils.Message(false, "User not found"))
		return
	}

	user.Delete()

	utils.Respond(w, utils.Message(true, "Account successfully deleted!"))
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	requesterId := r.Context().Value("user").(uint)

	if !models.IsAdmin(requesterId) {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	type request struct {
		UserID          uint   `json:"userId"`
		Accept          bool   `json:"accept"`
		RejectionReason string `json:"rejectionReason"`
	}

	payload := request{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid payload"))
		return
	}

	user := models.GetUser(payload.UserID, false)

	user.IsVerified = &payload.Accept
	user.RejectionReason = payload.RejectionReason

	models.GetDB().Save(user)

	utils.Respond(w, utils.Message(true, "success"))
}
