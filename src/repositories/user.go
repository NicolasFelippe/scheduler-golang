package repositories

import (
	"database/sql"
	"scheduler/src/models"
)

type users struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *users {
	return &users{db}
}

func (repository users) GetById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"select id, nome, empresa_id from users where id=?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	for rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.CompanyId,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
