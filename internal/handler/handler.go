package handler

import "net/http"

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetDMHistory(w http.ResponseWriter, r *http.Request)
	SendDM(w http.ResponseWriter, r *http.Request)
}

type Group interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	SendGroupMessage(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
	AddMember(w http.ResponseWriter, r *http.Request)
	RemoveMember(w http.ResponseWriter, r *http.Request)
	ListGroups(w http.ResponseWriter, r *http.Request)
	GetGroupHistory(w http.ResponseWriter, r *http.Request)
}

type Broadcast interface {
	BroadcastMessage(w http.ResponseWriter, r *http.Request)
}
