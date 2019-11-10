package controllers

import (
	"bytes"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"io"
	"io/ioutil"
	"net/http"
)

var GetUserImage = func(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint)
	account := models.GetUser(userId, false)
	email := account.Email

	var buf bytes.Buffer

	file, _, err := r.FormFile("file")
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
	}

	_, err = io.Copy(&buf, file)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
	}

	err = ioutil.WriteFile(fmt.Sprintf("images/users/%s.png", email), []byte(buf.String()), 0666)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
	}

	utils.Respond(w, utils.Message(true, "file received"))
	return
}
