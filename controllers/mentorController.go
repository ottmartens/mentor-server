package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
	"net/http"
)

func GetAvailableMentors(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)

	mentors := models.GetFreeMentors(userId)

	resp := utils.Message(true, "Success")
	resp["data"] = mentors

	utils.Respond(w, resp)

	return
}

func CreateFormingRequest(w http.ResponseWriter, r *http.Request) {
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
