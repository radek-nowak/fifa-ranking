package repository

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"github.com/radek-nowak/fifa-ranking/internal/team/model"
)

type TeamRepository struct {
	db *sql.DB
}

func NewRepository() (*TeamRepository, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}

	return &TeamRepository{db: db}, nil
}

func open() (*sql.DB, error) {
	connectionString := "postgres://postgres:pass@localhost:5432/?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *TeamRepository) GetTeams() []model.Team {
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
		err := rows.Scan(&team.ID, &team.Name, &team.Stars)
		if err != nil {
			return nil
		}
		teams = append(teams, team)
	}

	return teams
}

func (r *TeamRepository) GetTeamByName(name string) (model.Team, error) {
	db, err := open()
	defer db.Close()
	if err != nil {
		return model.Team{}, err
	}

	row := db.QueryRow("SELECT * FROM teams WHERE name=$1", name)

	var team model.Team
	err = row.Scan(&team.ID, &team.Name, &team.Stars)
	if err == sql.ErrNoRows {
		log.Fatal(err)
		return model.Team{}, errors.New("Could not find the team")
	}
	if err != nil {
		log.Fatal(err)
		return model.Team{}, err
	}

	return team, nil
}

func loadMockData(db *sql.DB) error {
	deleteQuery := `
    DELETE FROM teams;
    `

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS teams (
    id serial primary key,
    name varchar(100),
    stars integer
    );`

	insertDataQuery := `
    INSERT INTO teams (name, stars) values ('Liverpool', 5);
    INSERT INTO teams (name, stars) values ('FC Barcelona', 5);
    `

	_, err := db.Exec(deleteQuery)
	if err != nil {
		return err
	}

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	_, err = db.Exec(insertDataQuery)
	if err != nil {
		return err
	}

	return nil
}
