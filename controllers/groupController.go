package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

var GetGroups = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetGroups()
	resp := utils.Message(true, "success")
	resp["data"] = data

	utils.Respond(w, resp)
}

var CreateGroupDirectly = func(w http.ResponseWriter, r *http.Request) {

	type requestModel struct {
		Title       string `json:"title"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
		Mentors     []uint `json:"mentors"`
	}

	request := &requestModel{}
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
