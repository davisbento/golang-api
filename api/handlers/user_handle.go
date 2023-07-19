package handlers

import (
	"davisbento/golang-api/db/entity"
	"davisbento/golang-api/db/repository"
	"davisbento/golang-api/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (handler *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := handler.repo.FindAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.repo.FindById(id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := entity.NewUser()

	// transform the json into a user struct
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = handler.repo.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
