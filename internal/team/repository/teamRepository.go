package repository

import (
	"database/sql"

	"github.com/radek-nowak/fifa-ranking/internal/team/model"
)

type TeamRepository struct {
	db *sql.DB
}

func open() (*sql.DB, error) {
	connectionString := "postgres://postgres:pass@localhost:5432/testdb?sslmode=silent"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetTeams() []model.Team {
	db, err := open()
	if err != nil {
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		return nil
	}

	teams := []model.Team{}
	for rows.Next() {
		var team model.Team
		err := rows.Scan(&team.ID, &team.Namne, &team.Stars)
		if err != nil {
			return nil
		}
		teams = append(teams, team)
	}

	return teams
}
