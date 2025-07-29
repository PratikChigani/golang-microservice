// internal/handlers/user_handler.go
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pratikchigani/user-service/internal/models"
	"github.com/pratikchigani/user-service/internal/utils"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := models.GetUserByID(id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := models.CreateUser(u); err != nil {
		utils.RespondError(w, http.StatusConflict, "User already exists or invalid input")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedUser models.User

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.ValidateUser(updatedUser); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.UpdateUser(id, updatedUser); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "User not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to user book")
		}
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := models.DeleteUser(id); err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusNoContent, nil)
}
