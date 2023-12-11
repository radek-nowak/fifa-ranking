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

func (s *TeamService) GetTeam(name string) (model.Team, error) {
	team, err := s.repository.GetTeamByName(name)
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (s *TeamService) SaveNewTeam(team model.Team) error {
	err := s.repository.Save(team)
	if err != nil {
		return err
	}
	return nil
}
