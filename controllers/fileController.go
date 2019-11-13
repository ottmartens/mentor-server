package controllers

import (
	"bytes"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var GetUserImage = func(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)
	account := models.GetUser(userId, false)
	email := account.Email

	var buf bytes.Buffer

	file, _, err := r.FormFile("file")
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	_, err = io.Copy(&buf, file)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("images/users/%s.png", email), []byte(buf.String()), 0666)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	account.ImageUrl = fmt.Sprintf("/api/images/users/%s.png", email)
	models.GetDB().Save(account)

	resp := utils.Message(true, "file received")

	type ResponseData struct {
		ImageUrl string `json:"imageUrl"`
	}
	resp["data"] = ResponseData{
		ImageUrl: account.ImageUrl,
	}
	time.Sleep(5 * time.Second)

	utils.Respond(w, resp)
	return
}
