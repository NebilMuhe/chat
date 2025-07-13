package dto

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignUP struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SignUP) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.FullName, validation.Required.Error("full name is required")),
		validation.Field(&s.Email, validation.Required.Error("email is required"), is.Email),
		validation.Field(&s.Password, validation.Required.Error("password is required")),
	)
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required.Error("email is reuired")),
		validation.Field(&l.Password, validation.Required.Error("password is required")),
	)
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type DirectMessage struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}
