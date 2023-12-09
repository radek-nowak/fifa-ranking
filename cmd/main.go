package main

import (
	"github.com/gin-gonic/gin"

	controller "github.com/radek-nowak/fifa-ranking/internal/team/controller"
	"github.com/radek-nowak/fifa-ranking/internal/team/repository"
	"github.com/radek-nowak/fifa-ranking/internal/team/service"
)

func main() {
	r, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}

	s := service.NewService(r)
	con := controller.NewController(s)

	router := gin.Default()
	router.GET("/teams", con.GetTeams)
	router.GET("/teams/:name", con.GetTeamByName)
	//
	// router.POST("/teams", con.AddTeam)

	router.Run("localhost:9090")
}
