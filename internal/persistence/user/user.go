package user

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/persistence"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type user struct {
	redisClient *redis.Client
}

func InitPersistence(redisClient *redis.Client) persistence.User {
	return &user{
		redisClient: redisClient,
	}
}

func (u *user) CreateUser(ctx context.Context, user dto.SignUP) error {
	userMap := map[string]any{
		"full_name": user.FullName,
		"email":     user.Email,
		"password":  user.Password,
	}

	if err := u.redisClient.HSet(ctx, user.Email, userMap).Err(); err != nil {
		return err
	}

	return nil
}

func (u *user) UserExists(ctx context.Context, email string) error {
	exists, err := u.redisClient.Exists(ctx, email).Result()
	if err != nil {
		return err
	}

	if exists > 0 {
		return fmt.Errorf("user already exits")
	}

	return nil
}

func (u *user) GetUser(ctx context.Context, email string) (dto.SignUP, error) {
	data, err := u.redisClient.HGetAll(ctx, email).Result()
	if err != nil {
		return dto.SignUP{}, err
	}

	if len(data) == 0 {
		return dto.SignUP{}, fmt.Errorf("user not found")
	}

	return dto.SignUP{
		FullName: data["full_name"],
		Email:    data["email"],
		Password: data["password"],
	}, nil
}

func (u *user) SendDM(ctx context.Context, message dto.DirectMessage) error {
	message.Timestamp = time.Now().UTC()

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	chatKey := fmt.Sprintf("dm:%s:%s", message.From, message.To)
	if message.From > message.To {
		chatKey = fmt.Sprintf("dm:%s:%s", message.To, message.From)
	}

	if err := u.redisClient.RPush(ctx, chatKey, data).Err(); err != nil {
		return err
	}

	pubChannel := "dm:" + message.To
	return u.redisClient.Publish(ctx, pubChannel, data).Err()
}

func (u *user) GetDMHistory(ctx context.Context, user1, user2 string) ([]dto.DirectMessage, error) {
	key := fmt.Sprintf("dm:%s:%s", user1, user2)
	if user1 > user2 {
		key = fmt.Sprintf("dm:%s:%s", user2, user1)
	}
	messages, err := u.redisClient.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var result []dto.DirectMessage
	for _, msg := range messages {
		var dm dto.DirectMessage
		if err := json.Unmarshal([]byte(msg), &dm); err == nil {
			result = append(result, dm)
		}
	}
	return result, nil
}
