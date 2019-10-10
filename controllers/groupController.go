package controllers

import (
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
