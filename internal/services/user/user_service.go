package user

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/helpers"
	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/interfaces"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	UserRepo interfaces.IUserRepository
	Logger   *logrus.Logger
	Redis    *redis.Client
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		res dto.LoginResponse
		now = time.Now()
	)

	userData, err := s.UserRepo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		s.Logger.Error("service::Login - Failed to find user by username : ", err)
		return res, err
	}

	if !helpers.ComparePassword(userData.Password, req.Password) {
		s.Logger.Error("service::Login - Password not match")
		return res, errors.New(constants.ErrUsernameOrPasswordIsIncorrect)
	}

	token, err := helpers.GenerateToken(ctx, userData.ID.String(), userData.Username, userData.FullName, constants.TokenTypeAccess, now)
	if err != nil {
		s.Logger.Error("service::Login - Failed to generate token : ", err)
		return res, errors.New(constants.ErrFailedGenerateToken)
	}

	refreshToken, err := helpers.GenerateToken(ctx, userData.ID.String(), userData.Username, userData.FullName, constants.RefreshTokenAccess, now)
	if err != nil {
		s.Logger.Error("service::Login - Failed to generate refresh token : ", err)
		return res, errors.New(constants.ErrFailedGenerateRefreshToken)
	}

	res.UserID = userData.ID.String()
	res.Username = userData.Username
	res.FullName = userData.FullName
	res.Token = token
	res.RefreshToken = refreshToken

	return res, nil
}

func (s *UserService) Logout(ctx context.Context, token string) error {
	err := s.Redis.Del(ctx, token).Err()
	if err != nil {
		s.Logger.Error("service::Logout - Failed to delete token from redis : ", err)
		return err
	}

	return nil
}

func (s *UserService) GetUserProfile(ctx context.Context, id string) (*dto.GetUserProfileResponse, error) {
	userData, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		s.Logger.Error("service::GetUserProfile - Failed to find user by id : ", err)
		return nil, err
	}

	return &dto.GetUserProfileResponse{
		ID:          userData.ID.String(),
		Username:    userData.Username,
		FullName:    userData.FullName,
		Email:       userData.Email,
		Nim:         userData.Nim,
		PhoneNumber: userData.PhoneNumber,
	}, nil
}
