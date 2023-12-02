package main

import (
	"github.com/gin-gonic/gin"

	"ranking/cmd/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/teams", controllers.GetTeams)
	router.GET("/teams/:name", controllers.GetTeamByName)
	router.Run("localhost:9090")
}
