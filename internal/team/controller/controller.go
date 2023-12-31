package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/radek-nowak/fifa-ranking/internal/team/model"
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

func (t *TeamsController) GetTeamByName(ctx *gin.Context) {
	name := ctx.Param("name")
	team, err := t.service.GetTeam(name)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "team with name " + name + " not found"},
		)
	}
	ctx.IndentedJSON(http.StatusOK, team)
}

func (t *TeamsController) AddTeam(ctx *gin.Context) {
	var team model.Team

	if err := ctx.BindJSON(&team); err != nil {
		ctx.IndentedJSON(
			http.StatusEarlyHints,
			gin.H{"message": "Error occured when creating a new team"},
		)
		return
	}

	t.service.SaveNewTeam(team)
	ctx.IndentedJSON(http.StatusCreated, team)
}

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
