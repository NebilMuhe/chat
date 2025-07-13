package group

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/handler"
	"chat_system/internal/module"
	"encoding/json"
	"net/http"
)

type group struct {
	groupModule module.Group
}

func InitHandler(groupModule module.Group) handler.Group {
	return &group{
		groupModule: groupModule,
	}
}

func (g *group) AddMember(w http.ResponseWriter, r *http.Request) {
	var member dto.AddMemeber

	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	owner_id, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := g.groupModule.AddMember(r.Context(),
		member.GroupName, owner_id, member.UserID); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user added successfully"))
}

func (g *group) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group dto.CreateGroup

	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := group.Validate(); err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	if err := g.groupModule.CreateGroup(r.Context(), group); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("registered succesfully"))
}

func (g *group) GetGroup(w http.ResponseWriter, r *http.Request) {
	ownerID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := r.URL.Query().Get("group_name")
	if groupName == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	group, err := g.groupModule.GetGroup(r.Context(), ownerID, groupName)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

func (g *group) GetGroupHistory(w http.ResponseWriter, r *http.Request) {
	ownerID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := r.URL.Query().Get("group_name")
	if groupName == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	history, err := g.groupModule.GetGroupHistory(r.Context(), ownerID, groupName)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}

func (g *group) ListGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := g.groupModule.ListGroups(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

func (g *group) RemoveMember(w http.ResponseWriter, r *http.Request) {
	var member dto.AddMemeber

	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	ownerID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := g.groupModule.RemoveMember(
		r.Context(), ownerID, member.GroupName, member.UserID); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("removed successfully"))
}

func (g *group) SendGroupMessage(w http.ResponseWriter, r *http.Request) {
	var message dto.GroupMessage
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := r.URL.Query().Get("group_name")
	if groupName == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := g.groupModule.SendGroupMessage(r.Context(),
		userID, groupName, message); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("sent successfully"))
}
