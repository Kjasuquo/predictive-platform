package api

import (
	"net/http"
	"os"
	"predictive-platform/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"

	"predictive-platform/internal/domain/dto"
	"predictive-platform/internal/domain/services"
	"predictive-platform/pkg/config"
	"predictive-platform/pkg/web"
)

type Handler struct {
	Config      *config.Config
	UserService services.UserService
}

func NewHandler(cfg *config.Config, userService services.UserService) *Handler {
	return &Handler{
		Config:      cfg,
		UserService: userService,
	}
}

func (h *Handler) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "welcome to borderless money"})
	}
}

// AuthorizeUser authorizes a user request
func (h *Handler) AuthorizeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretkey := os.Getenv("JWT_SECRET")
		// Get token from header
		encodedToken := strings.TrimPrefix(ctx.GetHeader("Authorization"), config.BearerPrefix)
		if encodedToken == "" {
			web.Respond(ctx, http.StatusUnauthorized, nil, "invalid token", []string{"invalid token"})
			ctx.Abort()
			return
		}
		accessClaims, err := jwt.ValidateAndGetClaims(encodedToken, secretkey)
		if err != nil {
			web.Respond(ctx, http.StatusUnauthorized, nil, "unable to validate token", []string{err.Error()})
			ctx.Abort()
			return
		}

		email, ok := accessClaims["email"].(string)
		if !ok {
			web.Respond(ctx, http.StatusUnauthorized, nil, "unauthorized user", []string{"unauthorized user"})
			ctx.Abort()
			return
		}

		var user *dto.User
		if user, err = h.UserService.GetUserByEmail(ctx, email); err != nil {
			web.Respond(ctx, http.StatusUnauthorized, nil, "unable to retrieve user", []string{err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("access_token", encodedToken)
		ctx.Set("user", user)

		ctx.Next()
	}
}
