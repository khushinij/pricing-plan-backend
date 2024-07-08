package main

import (
	"HackOn/controllers"
	"HackOn/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	controllers.RegisterPricingRoutes(r)
	controllers.RegisterParameterRoutes(r)
	controllers.RegisterChatRoutes(r)
	controllers.SortedChatRoutes(r)

	r.Run(":8080")
}
