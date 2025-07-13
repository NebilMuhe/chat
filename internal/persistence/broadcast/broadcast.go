package broadcast

import (
	"chat_system/internal/persistence"
	"context"

	"github.com/redis/go-redis/v9"
)

type broadcast struct {
	redisClient *redis.Client
}

func InitPersistence(redisClient *redis.Client) persistence.Broadcast {
	return &broadcast{
		redisClient: redisClient,
	}
}

func (b *broadcast) BroadcastMessage(ctx context.Context, message string) error {
	key := "broadcast:messages"
	if err := b.redisClient.RPush(ctx, key, message).Err(); err != nil {
		return err
	}
	return nil
}
