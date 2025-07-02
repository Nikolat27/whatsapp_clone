package api

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"whatsapp_clone/internal/errors"
	authHelper "whatsapp_clone/internal/helpers/auth"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) GetUserChatsHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
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

func (app *Application) SearchUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
	}

	if err := jsonHelper.DeSerialize(r.Body, 10000, &input); err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := app.Models.User.GetUserInstanceByUsername(input.Username)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]any{
		"username":    input.Username,
		"user_id":     user.Id,
		"profile_url": user.ProfilePicUrl,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		errors.ServerErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *Application) SearchUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserId string `json:"user_id"`
	}

	if err := jsonHelper.DeSerialize(r.Body, 10000, &input); err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userObjectId, err := convertHelper.StringToObjectId(input.UserId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := app.Models.User.GetUserInstanceById(userObjectId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]any{
		"username":    user.Username,
		"user_id":     input.UserId,
		"profile_url": user.ProfilePicUrl,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		errors.ServerErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *Application) UploadProfilePicHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with 10MB max memory
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from posted form-data ("profile_picture")
	file, fileHeader, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Sanitize the file name to prevent directory traversal attacks
	filename := filepath.Base(fileHeader.Filename)

	// Ensure the upload directory exists (relative to your app working dir)
	uploadDir := "./uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to create upload directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the file to disk
	fullPath := filepath.Join(uploadDir, filename)
	fmt.Println(fullPath)

	err = os.WriteFile(fullPath, fileBytes, 0644)
	if err != nil {
		http.Error(w, "Unable to save the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := app.GetUserId(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	updates := bson.M{
		"profile_url": fullPath,
	}

	err = app.Models.User.UpdateUserInstance(userId, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Application) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if userId == primitive.NilObjectID {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "user id does not exist")
		return
	}

	user, err := app.Models.User.GetUserInstanceById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if user == nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "user does not exist")
		return
	}

	data := map[string]any{
		"username":    user.Username,
		"user_id":     user.Id,
		"about":       user.About,
		"name":        user.Name,
		"profile_url": user.ProfilePicUrl,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		errors.ServerErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (app *Application) GetUserId(r *http.Request) (primitive.ObjectID, error) {
	token, err := authHelper.GetUserAuthToken(r)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if token == nil {
		return primitive.NilObjectID, nil
	}

	userId := app.Rli.Get(string(token)).Val()

	userObjectId, err := convertHelper.StringToObjectId(userId)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userObjectId, nil
}
