package web

import (
	"log"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data      any      `json:"data"`
	Message   string   `json:"message"`
	Errors    []string `json:"errors,omitempty"`
	Status    string   `json:"status"`
	Timestamp string   `json:"timestamp"`
} // @name Response

func Respond(
	c *gin.Context, statusCode int, data any, message string, errs []string) {
	if c.Request.Context().Err() != nil {
		// probably client cancelled the request context
		// or context deadline reached
		return
	}

	if statusCode >= http.StatusBadRequest {
		log.Printf("an error ocurred: %v - %+v", message, errs)
	}

	responseData := Response{
		Message:   message,
		Data:      data,
		Errors:    errs,
		Status:    http.StatusText(statusCode),
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(statusCode, responseData)
}
