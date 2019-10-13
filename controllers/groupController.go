package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
	"net/http"
)

var GetGroups = func(w http.ResponseWriter, r *http.Request) {

	groups := models.GetGroups()
	resp := utils.Message(true, "success")

	resp["data"] = groups

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

var RequestGroupJoining = func(w http.ResponseWriter, r *http.Request) {

	type payload struct {
		GroupId uint `json:"groupId"`
		UserId  uint `json:"userId"`
	}

	request := &payload{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	group := models.GetGroup(request.GroupId)
	if group == nil {
		utils.Respond(w, utils.Message(false, "Invalid group id"))
		return
	}

	user := models.GetUser(request.UserId)
	if user == nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}
	if user.Role != enums.UserTypes.Mentee {
		utils.Respond(w, utils.Message(false, "User is not a mentee"))
		return
	}
	if user.GroupId != 0 {
		utils.Respond(w, utils.Message(false, "User already belongs to a group"))
		return
	}

	resp := models.CreateRequest(enums.RequestTypes.JoinGroup, request.UserId, request.GroupId)

	utils.Respond(w, resp)
}
