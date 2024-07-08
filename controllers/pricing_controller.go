package controllers

import (
	"HackOn/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PC struct {
	MerchantID         string `json:"MerchantId"`
	Round              string `json:"round"`
	Volume             string `json:"volume"`
	ServicesNegotiated string `json:"servicesNegotiated"`
	Competition        string `json:"competition"`
}

func RegisterPricingRoutes(r *gin.Engine) {
	r.POST("/pricing/", GetPricing)
}

func GetPricing(c *gin.Context) {
	var pc PC
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	negotiation, err := services.CalculatePricing(pc.MerchantID, pc.Round, pc.Volume, pc.ServicesNegotiated, pc.Competition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"negotiation": negotiation})
}
