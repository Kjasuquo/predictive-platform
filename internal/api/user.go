package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"predictive-platform/internal/domain/dto"
	"predictive-platform/pkg/web"
)

func (h *Handler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.User
		if err := ctx.ShouldBindJSON(&request); err != nil {
			web.Respond(ctx, http.StatusBadRequest, nil, "bad payload", []string{err.Error()})
			return
		}

		if !request.IsValid() {
			web.Respond(ctx, http.StatusBadRequest, nil, "invalid payload", []string{"invalid payload"})
			return
		}

		userID, token, err := h.UserService.UserSignUp(ctx, &request)
		if err != nil {
			web.Respond(ctx, http.StatusInternalServerError, nil, "Unable to register new user", []string{err.Error()})
			return
		}

		web.Respond(ctx, http.StatusCreated, gin.H{
			"user_id": fmt.Sprintf("%v", userID),
			"token":   fmt.Sprintf("%v", token),
		}, "successful", nil)
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request dto.Login

		if err := ctx.ShouldBindJSON(&request); err != nil {
			web.Respond(ctx, http.StatusBadRequest, nil, "bad payload", []string{err.Error()})
			return
		}

		user, token, err := h.UserService.UserLogin(ctx, &request)
		if err != nil {
			web.Respond(ctx, http.StatusBadRequest, nil, "bad payload", []string{err.Error()})
			return
		}

		web.Respond(ctx, http.StatusOK, gin.H{
			"user":  fmt.Sprintf("%v", user),
			"token": fmt.Sprintf("%v", token),
		}, "successful", nil)
	}
}
