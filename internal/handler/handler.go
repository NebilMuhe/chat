package handler

import "net/http"

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetDMHistory(w http.ResponseWriter, r *http.Request)
	SendDM(w http.ResponseWriter, r *http.Request)
}
