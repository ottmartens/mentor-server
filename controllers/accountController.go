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

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	requesterId := r.Context().Value("user").(uint)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}

	if !isAdmin(r) && uint(userId) != requesterId {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	user := models.GetUser(uint(userId), false)

	if user == nil {
		utils.Respond(w, utils.Message(false, "User not found"))
		return
	}

	utils.Respond(w, utils.Message(true, "Account successfully deleted!"))
}
