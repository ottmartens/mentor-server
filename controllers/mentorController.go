package controllers

import (
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func GetAvailableMentors(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)

	fmt.Println(userId)

	mentors := models.GetFreeMentors(userId)

	resp := utils.Message(true, "Success")
	resp["data"] = mentors

	utils.Respond(w, resp)

	return
}
