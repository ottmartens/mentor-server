package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
	"strconv"
)

var GetGroups = func(w http.ResponseWriter, r *http.Request) {
	groups := models.GetGroups()
	resp := utils.Message(true, "success")

	resp["data"] = groups

	utils.Respond(w, resp)
}

var GetGroupDetails = func(w http.ResponseWriter, r *http.Request) {
	groupId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid group id"))
		return
	}

	data := models.GetGroupDetails(uint(groupId))
	resp := utils.Message(true, "Success")
	resp["data"] = data

	utils.Respond(w, resp)
}

var GetUsersGroup = func(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)

	user := models.GetUser(userId, false)

	if user.GroupId == nil {
		utils.Respond(w, utils.Message(false, "You do not belong to a group"))
		return
	}

	data := models.GetGroupDetails(*user.GroupId)

	resp := utils.Message(true, "Success")

	resp["data"] = data

	utils.Respond(w, resp)
}

var CreateGroupDirectly = func(w http.ResponseWriter, r *http.Request) {

	type payload struct {
		Title       string `json:"title"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
		Mentors     []uint `json:"mentors"`
	}

	request := &payload{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	group := models.Group{
		Title:       request.Title,
		Tagline:     request.Tagline,
		Description: request.Description,
	}
	mentors := request.Mentors

	resp := group.Create(mentors)
	utils.Respond(w, resp)
}
