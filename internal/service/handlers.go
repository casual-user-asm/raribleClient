package service

import (
	"net/http"
	"time"

	"github.com/casual-user-asm/raribleClient/internal/client"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/ownership/:id", OwnershipHandler)
	router.POST("/traits", TraitsHandler)
	router.Run(":8080")
}

func OwnershipHandler(c *gin.Context) {
	ownershipID := c.Param("id")

	clientForFunc := &http.Client{Timeout: 10 * time.Second}

	ownership, err := client.RetrieveOwnershipByID(clientForFunc, ownershipID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ownership)
}

func TraitsHandler(c *gin.Context) {
	var req client.TraitRarityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	clientForFunc := &http.Client{Timeout: 10 * time.Second}

	traitsRarity, err := client.RetrieveTraitsRarity(clientForFunc, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, traitsRarity)
}
