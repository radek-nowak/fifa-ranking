package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/radek-nowak/fifa-ranking/internal/team/service"
)

type TeamsController struct {
	service *service.TeamService
}

func NewController(teamService *service.TeamService) *TeamsController {
	return &TeamsController{service: teamService}
}

func (t *TeamsController) GetTeams(ctx *gin.Context) {
	teams := t.service.GetTeams()
	ctx.IndentedJSON(http.StatusOK, teams)
}

//
// func GetTeamByName(ctx *gin.Context) {
// 	name := ctx.Param("name")
//
// 	for _, t := range teams {
// 		if t.Namne == name {
// 			ctx.IndentedJSON(http.StatusOK, t)
// 			return
// 		}
// 	}
//
// 	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "team with name " + name + " not found"})
// }
//
// func AddTeam(ctx *gin.Context) {
// 	var newTeam model.Team
//
// 	if err := ctx.BindJSON(&newTeam); err != nil {
// 		return
// 	}
//
// 	teams = append(teams, newTeam)
// 	ctx.IndentedJSON(http.StatusAccepted, newTeam)}
