package module

import (
	"chat_system/internal/constant/model/dto"
	"context"
)

type User interface {
	CreateUser(ctx context.Context, user dto.SignUP) error
	LoginUser(ctx context.Context, user dto.Login) (dto.LoginResponse, error)
	SendDM(ctx context.Context, message dto.DirectMessage) error
	GetDMHistory(ctx context.Context, user1, user2 string) ([]dto.DirectMessage, error)
}
