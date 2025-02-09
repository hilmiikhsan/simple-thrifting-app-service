package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/models"
)

type IUserRepository interface {
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	FindUserByID(ctx context.Context, id string) (*models.User, error)
}

type IUserService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (dto.LoginResponse, error)
	Logout(ctx context.Context, token string) error
	GetUserProfile(ctx context.Context, id string) (*dto.GetUserProfileResponse, error)
}

type IUserHandler interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	GetUserProfile(*gin.Context)
}
