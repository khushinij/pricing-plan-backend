package controllers

import (
	"HackOn/models"
	"HackOn/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterParameterRoutes(r *gin.Engine) {
	r.POST("/parameter", CreateParameter)
}

func CreateParameter(c *gin.Context) {
	var param models.Parameter
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(param)

	if err := services.CreateParameter(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, param)
}
