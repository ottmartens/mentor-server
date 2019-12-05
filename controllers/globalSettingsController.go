package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func GetGlobalSettings(w http.ResponseWriter, _ *http.Request) {

	settings := models.GetGlobalSettings()

	resp := utils.Message(true, "success")
	resp["data"] = settings

	utils.Respond(w, resp)
}

func SetGlobalSettings(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	if !models.IsAdmin(userId) {
		utils.Respond(w, utils.Message(false, "not permitted"))
		return
	}

	request := map[string]interface{}{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		fmt.Println(err.Error())
		utils.Respond(w, utils.Message(false, "invalid request"))
	}

	resp := models.SaveGlobalSettings(request)

	utils.Respond(w, resp)
}
