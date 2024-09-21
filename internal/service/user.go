package service

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/go-tibia-info-back/internal/database"
	"github.com/gaspartv/go-tibia-info-back/internal/entity"
	handlerError "github.com/gaspartv/go-tibia-info-back/internal/handleError"
	"github.com/go-chi/chi/v5"
)

type UserService struct {
	db database.UserDB
}

func NewUserService(db database.UserDB) *UserService {
	return &UserService{
		db: db,
	}
}

func (service *UserService) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.FirstName == "" {
		handlerError.Exec(w, "first name is required", http.StatusBadRequest)
		return
	}

	if user.LastName == "" {
		handlerError.Exec(w, "last name is required", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		handlerError.Exec(w, "email is required", http.StatusBadRequest)
		return
	}

	if user.NationalId == "" {
		handlerError.Exec(w, "national id is required", http.StatusBadRequest)
		return
	}

	if user.Telephone == "" {
		handlerError.Exec(w, "telephone id is required", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		handlerError.Exec(w, "telephone id is required", http.StatusBadRequest)
		return
	}

	if user.BirthDate == "" {
		handlerError.Exec(w, "birth date id is required", http.StatusBadRequest)
		return
	}

	if user.Language == "" {
		handlerError.Exec(w, "language id is required", http.StatusBadRequest)
		return
	}

	if user.DarkMode != true && user.DarkMode != false {
		handlerError.Exec(w, "dark mode id is required", http.StatusBadRequest)
		return
	}

	if len(user.Permissions) == 0 {
		handlerError.Exec(w, "permissions id is required", http.StatusBadRequest)
		return
	}

	isNotUnique := service.db.VerifyUnique(user.Email)
	if isNotUnique {
		handlerError.Exec(w, "user already exists", http.StatusConflict)
		return
	}

	s := entity.NewUser(entity.CreateUserDTO{
		Email:       user.Email,
		Password:    user.Password,
		Code:        "951651",
		Type:        user.Type,
		Police:      user.Police,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		NationalId:  user.NationalId,
		Telephone:   user.Telephone,
		BirthDate:   user.BirthDate,
		Language:    user.Language,
		DarkMode:    user.DarkMode,
		Permissions: user.Permissions,
	})
	result, err := service.db.Create(s)

	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *UserService) Update(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		handlerError.Exec(w, "email is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.db.Get(); err != nil {
		handlerError.Exec(w, "user not found", http.StatusNotFound)
		return
	}

	result, err := handler.db.Update(id, &user)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *UserService) Get(w http.ResponseWriter, r *http.Request) {
	result, err := handler.db.Get()
	if err != nil {
		handlerError.Exec(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.db.Get(); err != nil {
		handlerError.Exec(w, "user not found", http.StatusNotFound)
		return
	}

	result, err := handler.db.Delete(id)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
