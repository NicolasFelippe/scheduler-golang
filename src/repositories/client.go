package repositories

import (
	"database/sql"
	"scheduler/src/models"
)

type clients struct {
	db *sql.DB
}

func NewRepositoryClient(db *sql.DB) *clients {
	return &clients{db}
}

func (repository clients) GetById(id uint64) (models.Client, error) {
	rows, err := repository.db.Query(
		"select c.id, c.name from clients c where id=?",
		id,
	)
	if err != nil {
		return models.Client{}, err
	}
	defer rows.Close()

	var client models.Client

	for rows.Next() {
		if err = rows.Scan(
			&client.ID,
			&client.Name,
		); err != nil {
			return models.Client{}, err
		}
	}

	return client, nil
}
