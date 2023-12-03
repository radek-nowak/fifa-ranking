package main

import (
	"github.com/gin-gonic/gin"

	controller "github.com/radek-nowak/fifa-ranking/internal/team/controller"
)

func main() {
	router := gin.Default()

	router.GET("/teams", controller.GetTeams)
	router.GET("/teams/:name", controller.GetTeamByName)

	router.POST("/teams", controller.AddTeam)

	router.Run("localhost:9090")
}
