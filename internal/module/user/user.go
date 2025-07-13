package user

import (
	"chat_system/internal/constant/model/dto"
	"chat_system/internal/module"
	"chat_system/internal/persistence"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	userPersistence persistence.User
	secretKey       string
}



func InitUserModule(userPersistence persistence.User, secretKey string) module.User {
	return &user{
		userPersistence: userPersistence,
		secretKey:       secretKey,
	}
}

func (u *user) CreateUser(ctx context.Context, user dto.SignUP) error {
	if err := u.userPersistence.UserExists(ctx, user.Email); err != nil {
		return err
	}

	password, err := u.hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password

	return u.userPersistence.CreateUser(ctx, user)
}

func (u *user) LoginUser(ctx context.Context, user dto.Login) (dto.LoginResponse, error) {
	userRes, err := u.userPersistence.GetUser(ctx, user.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if err := u.comparePassword(user.Password, userRes.Password); err != nil {
		return dto.LoginResponse{}, err
	}

	accessToken, err := u.generateAccessToken(user.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken: accessToken,
		ExpiredAt:   time.Now().Add(15 * time.Minute),
	}, nil
}

func (u *user) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *user) comparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *user) generateAccessToken(email string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "chat-app",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *user) GetDMHistory(ctx context.Context, user1 string, user2 string) ([]dto.DirectMessage, error) {
	return u.userPersistence.GetDMHistory(ctx, user1, user2)
}

func (u *user) SendDM(ctx context.Context, message dto.DirectMessage) error {
	return u.userPersistence.SendDM(ctx,message)
}
