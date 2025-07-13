package user

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/handler"
	"chat_system/internal/module"
	"encoding/json"
	"net/http"
)

type user struct {
	userModule module.User
}

func InitUserHandler(userModule module.User) handler.UserHandler {
	return &user{
		userModule: userModule,
	}
}

func (u *user) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.SignUP

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := u.userModule.CreateUser(r.Context(), user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *user) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user dto.Login

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	res, err := u.userModule.LoginUser(r.Context(), user)
	if err != nil {
		http.Error(w, "failed to log in user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
	}
}

func (u *user) GetDMHistory(w http.ResponseWriter, r *http.Request) {
	user1 := r.Context().Value("user").(string)
	user2 := r.URL.Query().Get("user")
	if user2 == "" {
		http.Error(w, "missing user", http.StatusBadRequest)
		return
	}
	messages, err := u.userModule.GetDMHistory(r.Context(), user1, user2)
	if err != nil {
		http.Error(w, "failed to get messages", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (u *user) SendDM(w http.ResponseWriter, r *http.Request) {
	var message dto.DirectMessage

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	if err := u.userModule.SendDM(r.Context(), message); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("message sent successfully")
}
