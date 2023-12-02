package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	ID    int32  `json:"id"`
	Namne string `json:"name"`
	Stars int    `json:"stars"`
}

var teams = []Team{
	{ID: 1, Namne: "Liverpool", Stars: 5},
	{ID: 2, Namne: "Manchester City", Stars: 5},
}

func GetTeams(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, teams)
}

func GetTeamByName(ctx *gin.Context) {
	name := ctx.Param("name")

	for _, t := range teams {
		if t.Namne == name {
			ctx.IndentedJSON(http.StatusOK, t)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "team with name " + name + " not found"})
}
