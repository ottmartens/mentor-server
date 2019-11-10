package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
	"net/http"
)

func RequestGroupJoining(w http.ResponseWriter, r *http.Request) {

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

	user := models.GetUser(request.UserId, true)
	if user == nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}
	if user.Role != enums.UserTypes.Mentee {
		utils.Respond(w, utils.Message(false, "User is not a mentee"))
		return
	}
	if user.GroupId != nil {
		utils.Respond(w, utils.Message(false, "User already belongs to a group"))
		return
	}
	resp := models.CreateRequest(enums.RequestTypes.JoinGroup, request.UserId, request.GroupId)

	utils.Respond(w, resp)
}

func HandleJoining(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	mentor := models.GetUser(userId, false)

	if mentor.GroupId == nil {
		utils.Respond(w, utils.Message(false, "You do not belong to a group!"))
		return
	}

	type payload struct {
		UserId uint `json:"userId"`
		Accept bool `json:"accept"`
	}
	request := &payload{
		Accept: true, // Accept request if no value present
	}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.HandleJoiningRequest(*mentor.GroupId, request.UserId, request.Accept)

	utils.Respond(w, resp)
}

func RequestGroupForming(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	type payload struct {
		UserId uint `json:"userId"`
	}
	request := &payload{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.CreateRequest(enums.RequestTypes.CreateGroup, userId, request.UserId)

	utils.Respond(w, resp)
	return
}

func HandleForming(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	type payload struct {
		UserId uint `json:"userId"`
		Accept bool `json:"accept"`
	}
	request := &payload{
		Accept: true, // Accept request if no value present
	}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.HandleFormingRequest(userId, request.UserId, request.Accept)

	utils.Respond(w, resp)
}
