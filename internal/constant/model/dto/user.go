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
	From      string    `json:"from"`
	To        string    `json:"to"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func (d DirectMessage) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.To, validation.Required),
		validation.Field(&d.Text, validation.Required),
	)
}

type Group struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	UserID    string `json:"user_id"`
}

type CreateGroup struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	UserID    string `json:"user_id"`
}

func (c CreateGroup) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.GroupName, validation.Required),
	)
}

type GroupMessage struct {
	Sender    string `json:"sender"`
	GroupID   string `json:"group_id"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func (g GroupMessage) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Content, validation.Required),
	)
}

type AddMemeber struct {
	GroupName string `json:"group_name"`
	UserID    string `json:"user_id"`
}

type BroadCast struct{
	Message string `json:"message"` 
}
