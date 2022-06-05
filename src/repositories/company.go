package repositories

import (
	"database/sql"
	"scheduler/src/models"
)

type companies struct {
	db *sql.DB
}

func NewRepositoryCompany(db *sql.DB) *companies {
	return &companies{db}
}

func (repository companies) GetById(id uint64) (models.Company, error) {
	rows, err := repository.db.Query(
		"select id, nome from empresas where id=?",
		id,
	)
	if err != nil {
		return models.Company{}, err
	}
	defer rows.Close()

	var company models.Company

	for rows.Next() {
		if err = rows.Scan(
			&company.ID,
			&company.Name,
		); err != nil {
			return models.Company{}, err
		}
	}

	return company, nil
}
