package persistence

import (
	"chat_system/internal/constant/model/dto"
	"context"
)

type User interface {
	CreateUser(ctx context.Context, user dto.SignUP) error
	UserExists(ctx context.Context,email string) error
	GetUser(ctx context.Context, email string) (dto.SignUP,error)
	SendDM(ctx context.Context, message dto.DirectMessage) error
	ReceiveDM(ctx context.Context, userID string)
	GetDMHistory(ctx context.Context, user1, user2 string) ([]dto.DirectMessage, error)
}
