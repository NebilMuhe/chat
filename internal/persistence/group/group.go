package group

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/persistence"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type group struct {
	redisClient *redis.Client
}

func InitPersistence(redisClient *redis.Client) persistence.Group {
	return &group{
		redisClient: redisClient,
	}
}

func (g *group) CreateGroup(ctx context.Context, group dto.CreateGroup) error {
	groupMap := map[string]any{
		"group_id":   group.GroupID,
		"group_name": group.GroupName,
		"user_id":    group.UserID,
	}

	if err := g.redisClient.HSet(ctx,
		fmt.Sprintf("%s:%s", group.UserID, group.GroupName), groupMap).Err(); err != nil {
		return err
	}

	return nil
}

func (g *group) GroupExists(ctx context.Context, userID, groupName string) error {
	exists, err := g.redisClient.Exists(ctx,
		fmt.Sprintf("%s:%s", userID, groupName)).Result()
	if err != nil {
		return err
	}

	if exists > 0 {
		return fmt.Errorf("group already exits")
	}

	return nil
}

func (g *group) AddMember(ctx context.Context, ownerID, groupName, userID string) error {
	if err := g.redisClient.SAdd(ctx,
		fmt.Sprintf("%s:%s", ownerID, groupName), userID).Err(); err != nil {
		return err
	}

	return nil
}

func (g *group) GetGroup(ctx context.Context, ownerID, groupName string) (*dto.Group, error) {
	groupMap, err := g.redisClient.HGetAll(ctx,
		fmt.Sprintf("%s:%s", ownerID, groupName)).Result()

	if err != nil {
		return nil, err
	}

	return &dto.Group{
		GroupID:   groupMap["group_id"],
		GroupName: groupMap["group_name"],
		UserID:    groupMap["user_id"],
	}, nil
}

func (g *group) GetGroupHistory(ctx context.Context, ownerID, groupName string) ([]dto.GroupMessage, error) {
	msgs, err := g.redisClient.LRange(ctx,
		fmt.Sprintf("%s:%s", ownerID, groupName), 0, -1).Result()

	if err != nil {
		return nil, err
	}

	var history []dto.GroupMessage

	for _, msg := range msgs {
		var m dto.GroupMessage
		if err := json.Unmarshal([]byte(msg), &m); err != nil {
			continue
		}
		history = append(history, m)
	}

	return history, nil
}

func (g *group) ListGroups(ctx context.Context) ([]dto.Group, error) {
	var groups []dto.Group

	iter := g.redisClient.Scan(ctx, 0, "*:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		groupMap, err := g.redisClient.HGetAll(ctx, key).Result()
		if err != nil || len(groupMap) == 0 {
			continue
		}
		groups = append(groups, dto.Group{
			GroupID:   groupMap["group_id"],
			GroupName: groupMap["group_name"],
			UserID:    groupMap["user_id"],
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *group) RemoveMember(ctx context.Context, ownerID, groupName string, userID string) error {

	if err := g.redisClient.SRem(ctx,
		fmt.Sprintf("%s:%s", ownerID, groupName), userID).Err(); err != nil {
		return err
	}
	return nil
}

func (g *group) SendGroupMessage(ctx context.Context, ownerID, groupName string, message dto.GroupMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := g.redisClient.RPush(ctx,
		fmt.Sprintf("%s:%s", ownerID, groupName), data).Err(); err != nil {
		return err
	}

	return nil
}
