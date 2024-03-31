package handlers

import (
	"bot/models"
	"bot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Assistant(c *gin.Context, clients map[string]chan []byte) {
	var message models.MessageData
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("id:", message.ID, "message:", message.Message)
	//client := clients["aaa"]
	//client <- []byte(message.Message)
	err := utils.Chat(message.Message, clients["aaa"])
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "Data sent to all clients")
}
