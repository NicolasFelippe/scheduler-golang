package client

import (
	"scheduler/src/db"
	"scheduler/src/models"
	"scheduler/src/repositories"
)

func GetById(id uint64) (models.Client, error) {

	db, err := db.Connect()
	if err != nil {
		return models.Client{}, err
	}

	defer db.Close()

	repository := repositories.NewRepositoryClient(db)
	client, err := repository.GetById(id)

	if err != nil {
		return models.Client{}, err
	}

	return client, nil
}
