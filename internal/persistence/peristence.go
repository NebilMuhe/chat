package persistence

import (
	"chat_system/internal/constant/model/dto"
	"context"
)

type User interface {
	CreateUser(ctx context.Context, user dto.SignUP) error
	UserExists(ctx context.Context, email string) error
	GetUser(ctx context.Context, email string) (dto.SignUP, error)
	SendDM(ctx context.Context, message dto.DirectMessage) error
	GetDMHistory(ctx context.Context, user1, user2 string) ([]dto.DirectMessage, error)
}

type Group interface {
	CreateGroup(ctx context.Context, group dto.CreateGroup) error
	SendGroupMessage(ctx context.Context,ownerID,groupName string,message dto.GroupMessage) error
	GetGroup(ctx context.Context, ownerID, groupName string) (*dto.Group, error)
	AddMember(ctx context.Context, groupID string, ownerID, userID string) error
	RemoveMember(ctx context.Context, ownerID, groupName string, userID string) error
	ListGroups(ctx context.Context) ([]dto.Group, error)
	GetGroupHistory(ctx context.Context, ownerID, groupName string) ([]dto.GroupMessage, error)
	GroupExists(ctx context.Context, userID, groupName string) error
}

type Broadcast interface {
	BroadcastMessage(ctx context.Context, message string) error
}
