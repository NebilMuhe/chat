package group

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/module"
	"chat_system/internal/persistence"
	"context"
)

type group struct {
	groupPersistence persistence.Group
}

func InitModule(groupPersistence persistence.Group) module.Group {
	return &group{
		groupPersistence: groupPersistence,
	}
}

func (g *group) AddMember(ctx context.Context, groupID string, ownerID string, userID string) error {
	return g.groupPersistence.AddMember(ctx, groupID, ownerID, userID)
}

func (g *group) CreateGroup(ctx context.Context, group dto.CreateGroup) error {
	if err := g.groupPersistence.GroupExists(ctx,
		group.UserID, group.GroupName); err != nil {
		return err
	}

	return g.groupPersistence.CreateGroup(ctx, group)
}

func (g *group) GetGroup(ctx context.Context, ownerID string, groupName string) (*dto.Group, error) {
	return g.groupPersistence.GetGroup(ctx, ownerID, groupName)
}

func (g *group) GetGroupHistory(ctx context.Context, ownerID string, groupName string) ([]dto.GroupMessage, error) {
	return g.groupPersistence.GetGroupHistory(ctx, ownerID, groupName)
}

func (g *group) ListGroups(ctx context.Context) ([]dto.Group, error) {
	return g.groupPersistence.ListGroups(ctx)
}


func (g *group) RemoveMember(ctx context.Context, ownerID string, groupName string, userID string) error {
	return g.groupPersistence.RemoveMember(ctx,ownerID,groupName,userID)
}


func (g *group) SendGroupMessage(ctx context.Context, ownerID string, groupName string, message dto.GroupMessage) error {
	return g.groupPersistence.SendGroupMessage(ctx,ownerID,groupName,message)
}
