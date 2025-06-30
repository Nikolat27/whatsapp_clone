package api

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"whatsapp_clone/internal/errors"
	authHelper "whatsapp_clone/internal/helpers/auth"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) GetUserChatsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserId string `json:"user_id"`
	}

	if err := jsonHelper.DeSerialize(r.Body, 10000, &input); err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := convertHelper.StringToObjectId(input.UserId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chats, err := app.Models.Chat.GetUserChats(userId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonHelper.WriteJSON(w, http.StatusOK, chats)
}

func (app *Application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		About    string `json:"about"`
	}

	err = jsonHelper.DeSerialize(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	updateDoc := bson.M{}
	
	if input.Username != "" {
		updateDoc["username"] = input.Username
	}
	
	if input.Name != "" {
		updateDoc["name"] = input.Name
	}
	
	if input.About != "" {
		updateDoc["about"] = input.About
	}

	if len(updateDoc) == 0 {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "No fields provided")
		return
	}

	err = app.Models.User.UpdateUserInstance(userId, updateDoc)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Application) SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
	}

	err := jsonHelper.DeSerialize(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
	}
}

func (app *Application) UploadProfilePicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("profilePic")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	fmt.Printf("File Size: %+v\n", fileHeader.Size)
	fmt.Printf("MIME Header: %+v\n", fileHeader.Header)

	// Read file data if needed (e.g., save to disk or upload to cloud)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	// Save the file, e.g., to disk
	err = os.WriteFile("./uploads/"+fileHeader.Filename, fileBytes, 0644)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Application) GetUserId(r *http.Request) (primitive.ObjectID, error) {
	token, err := authHelper.GetUserAuthToken(r)
	if err != nil {
		return primitive.NilObjectID, err
	}

	userId := app.Rli.Get(string(token)).Val()

	userObjectId, err := convertHelper.StringToObjectId(userId)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userObjectId, nil
}
