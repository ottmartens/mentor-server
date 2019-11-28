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

	imageId := utils.Uuid(12)
	imageUrl := fmt.Sprintf("images/users/%s.png", imageId)

	err = ioutil.WriteFile(imageUrl, []byte(buf.String()), 0666)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	account.ImageUrl = imageUrl
	models.GetDB().Save(account)

	resp := utils.Message(true, "file received")

	type ResponseData struct {
		ImageUrl string `json:"imageUrl"`
	}
	resp["data"] = ResponseData{
		ImageUrl: account.ImageUrl,
	}

	utils.Respond(w, resp)
	return
}
