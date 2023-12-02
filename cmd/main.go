package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type team struct {
	ID    int32  `json:"id"`
	Namne string `json:"name"`
	Stars int    `json:"stars"`
}

var teams = []team{
	{ID: 1, Namne: "Liverpool", Stars: 5},
	{ID: 2, Namne: "Manchester City", Stars: 5},
}

func getTeams(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, teams)
}

func main() {
	router := gin.Default()
	router.GET("/teams", getTeams)
	router.Run("localhost:9090")
}
