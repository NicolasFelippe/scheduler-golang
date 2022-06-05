package company

import (
	"scheduler/src/db"
	"scheduler/src/models"
	"scheduler/src/repositories"
)

func GetById(id uint64) (models.Company, error) {

	db, err := db.Connect()
	if err != nil {
		return models.Company{}, err
	}

	defer db.Close()

	repository := repositories.NewRepositoryCompany(db)
	company, err := repository.GetById(id)

	if err != nil {
		return models.Company{}, err
	}

	return company, nil
}
