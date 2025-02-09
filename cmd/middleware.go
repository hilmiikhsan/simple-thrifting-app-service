package cmd

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/helpers"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	token := helpers.ExtractBearerToken(authHeader)
	if token == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - invalid bearer token format")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorizationFormat))
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to validate token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - token is already expired")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenExpired))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, claims)

	ctx.Next()
}

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - authorization empty")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrAuthorizationIsEmpty))
		ctx.Abort()
		return
	}

	token := helpers.ExtractBearerToken(authHeader)
	if token == "" {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - invalid bearer token format")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrInvalidAuthorizationFormat))
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - failed to validate refresh token: ", err)
		ctx.JSON(http.StatusUnauthorized, helpers.Error(err))
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		helpers.Logger.Error("middleware::MiddlewareValidateAuth - refresh token is already expired")
		ctx.JSON(http.StatusUnauthorized, helpers.Error(constants.ErrTokenExpired))
		ctx.Abort()
		return
	}

	ctx.Set(constants.TokenTypeAccess, claims)

	ctx.Next()
}
