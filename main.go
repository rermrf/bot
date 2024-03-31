package main

import (
	"bot/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	r := gin.Default()

	clients := make(map[string]chan []byte)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin"},
		MaxAge:       12 * time.Hour,
	}))
	r.Use(func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		}
	})
	r.GET("/chats", func(c *gin.Context) {
		log.Info("GET /chats")
		handlers.Chats(c, clients)
	})
	r.POST("/assistant", func(c *gin.Context) {
		handlers.Assistant(c, clients)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
