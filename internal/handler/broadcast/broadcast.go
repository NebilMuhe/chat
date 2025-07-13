package broadcast

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/handler"
	"chat_system/internal/module"
	"encoding/json"
	"net/http"
)

type broadcast struct {
	broadcastModule module.Broadcast
}

func InitBroadcastHandler(broadcastModule module.Broadcast) handler.Broadcast {
	return &broadcast{
		broadcastModule: broadcastModule,
	}
}

func (b *broadcast) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
	var message dto.BroadCast

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := b.broadcastModule.BroadcastMessage(r.Context(),
		message.Message); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message sent"))
}
