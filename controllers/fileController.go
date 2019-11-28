package controllers

import (
	"bytes"
	"fmt"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	imagePath := fmt.Sprintf("images/users/%s.png", imageId)

	imageUrl := "/api/" + imagePath

	err = ioutil.WriteFile(imagePath, []byte(buf.String()), 0666)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	// Delete old file if present
	if account.ImageUrl != "" {
		oldImagePath := account.ImageUrl[5:len(account.ImageUrl)]
		fmt.Printf("Attempting to delete file at %s - ", oldImagePath)

		var _, err = os.Stat(oldImagePath)

		if err == nil {
			err = os.Remove(oldImagePath)

			if err == nil {
				fmt.Println("success!")
			} else {
				fmt.Println("unsuccessful")
			}
		} else if os.IsNotExist(err) {
			fmt.Println("file not found")
		} else {
			fmt.Println("error getting fileStat")
		}
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
