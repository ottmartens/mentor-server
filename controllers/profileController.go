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

func GetUserSelf(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	data := models.GetUser(userId, true)

	resp := utils.Message(true, "Success")
	resp["data"] = data

	utils.Respond(w, resp)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}

	user := models.GetUser(uint(userId), true)
	if user == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("Cannot find user with id %d", userId)))
		return
	}

	resp := utils.Message(true, "Success")

	resp["data"] = user.GetPublicInfo()

	utils.Respond(w, resp)
	return
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	type request struct {
		Name    string `json:"name"`
		Tagline string `json:"tagline"`
		Degree  string `json:"degree"`
		Year    string `json:"year"`
		Bio     string `json:"bio"`
	}

	req := request{}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
	}

	user := models.GetUser(userId, false)
	if user == nil {
		utils.Respond(w, utils.Message(false, "user not found"))
		return
	}

	if len(req.Name) > 0 {
		user.Name = req.Name
	}
	if len(req.Tagline) > 0 {
		user.Tagline = req.Tagline
	}
	if len(req.Degree) > 0 {
		user.Degree = req.Degree
	}
	if len(req.Year) > 0 {
		user.Year = req.Year
	}
	if len(req.Bio) > 0 {
		user.Bio = req.Bio
	}

	if user.IsVerified != nil && *user.IsVerified == false {
		user.IsVerified = nil
		user.RejectionReason = ""
	}

	models.GetDB().Save(user)

	resp := utils.Message(true, "Profile successfully edited")

	resp["data"] = struct {
		Name string `json:"name"`
	}{Name: user.Name}

	utils.Respond(w, resp)
}

func EditGroupProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	user := models.GetUser(userId, false)

	if user.GroupId == nil {
		utils.Respond(w, utils.Message(false, "You do not belong to a group!"))
		return
	}

	group := models.GetGroup(*user.GroupId)
	if group == nil {
		utils.Respond(w, utils.Message(false, "Group not found"))
		return
	}

	type request struct {
		Title       string `json:"title"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
	}

	profile := request{}
	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	if len(profile.Title) > 0 {
		group.Title = profile.Title
	}
	if len(profile.Description) > 0 {
		group.Description = profile.Description
	}
	if len(profile.Tagline) > 0 {
		group.Tagline = profile.Tagline
	}

	models.GetDB().Save(group)

	utils.Respond(w, utils.Message(true, "Profile successfully edited"))
}
