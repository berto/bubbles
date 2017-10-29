package queries

import (
	"database/sql"
)

type Team struct {
	TeamID   int    `json:"team_id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func GetTeams(db *sql.DB) ([]Team, error) {
	rows, err := db.Query(`SELECT * FROM team;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []Team
	for rows.Next() {
		team := Team{}
		rows.Scan(&team.TeamID, &team.Name, &team.ImageURL)
		teams = append(teams, team)
	}

	return teams, nil
}
