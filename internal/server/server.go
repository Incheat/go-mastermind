package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/incheat/go-mastermind/internal/todo"
)

func New() *gin.Engine {
	r := gin.New()

	// ---- Middleware ----
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(requestIDMiddleware())

	// ---- Routes ----

	// Health-check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API group
	api := r.Group("/api/v1")
	{
		todo.RegisterRoutes(api)
	}

	return r
}

// Simple request ID middleware
func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewString()
		c.Set("request_id", id)
		c.Writer.Header().Set("X-Request-ID", id)
		c.Next()
	}
}
