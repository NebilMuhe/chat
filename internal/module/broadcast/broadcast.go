package broadcast

import (
	"chat_system/internal/module"
	"chat_system/internal/persistence"
	"context"
)

type broadcast struct {
	broadcastPersistence persistence.Broadcast
}

func InitModule(broadcastPersistence persistence.Broadcast) module.Broadcast {
	return &broadcast{
		broadcastPersistence: broadcastPersistence,
	}
}

func (b *broadcast) BroadcastMessage(ctx context.Context, message string) error {
	return b.broadcastPersistence.BroadcastMessage(ctx, message)
}
