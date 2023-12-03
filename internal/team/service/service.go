package service

import (
	"github.com/radek-nowak/fifa-ranking/internal/team/model"
	"github.com/radek-nowak/fifa-ranking/internal/team/repository"
)

type TeamService struct {
	repository *repository.TeamRepository
}

func NewService(repository *repository.TeamRepository) *TeamService {
	return &TeamService{repository: repository}
}

func (s *TeamService) GetTeams() []model.Team {
	return s.repository.GetTeams()
}
