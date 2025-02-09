package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/helpers"
	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/interfaces"
	"github.com/hilmiikhsan/thrifting-app-service/internal/validator"
)

type UserHandler struct {
	UserService interfaces.IUserService
	Validator   *validator.Validator
}

func (api *UserHandler) Login(ctx *gin.Context) {
	var (
		req = new(dto.LoginRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	res, err := api.UserService.Login(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameOrPasswordIsIncorrect) {
			helpers.Logger.Error("handler::Login - Username or password is incorrect : ", err)
			ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrUsernameOrPasswordIsIncorrect))
			return
		}

		helpers.Logger.Error("handler::Login - Failed to login user : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *UserHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("handler::Logout - Authorization header is empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenIsEmpty))
		return
	}

	token := helpers.ExtractBearerToken(authHeader)

	err := api.UserService.Logout(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("handler::Logout - Failed to logout user : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(nil, ""))
}

func (api *UserHandler) GetUserProfile(ctx *gin.Context) {
	claimsValue, exists := ctx.Get(constants.TokenTypeAccess)
	if !exists {
		helpers.Logger.Error("handler::GetUserProfile - token claims not found")
		ctx.JSON(http.StatusUnauthorized, helpers.Error("token claims not found"))
		return
	}

	claims, ok := claimsValue.(*helpers.ClaimToken)
	if !ok {
		helpers.Logger.Error("handler::GetUserProfile - invalid token claims type")
		ctx.JSON(http.StatusUnauthorized, helpers.Error("invalid token claims"))
		return
	}

	id := claims.UserID

	res, err := api.UserService.GetUserProfile(ctx.Request.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUserNotFound) {
			helpers.Logger.Error("handler::GetUserProfile - User not found : ", err)
			ctx.JSON(http.StatusNotFound, helpers.Error(constants.ErrUserNotFound))
			return
		}

		helpers.Logger.Error("handler::GetUserProfile - Failed to get user profile : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}
