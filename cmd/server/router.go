package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"predictive-platform/internal/api"
)

func DefineRoutes(handler *api.Handler) *gin.Engine {
	log.Println("Routes defined")

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	{
		router.GET("/", handler.Home())
	}

	r := router.Group("/api/v1")

	{
		r.POST("/register", handler.Register())
		r.POST("/login", handler.Login())
	}

	return router
}
