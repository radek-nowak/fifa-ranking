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
