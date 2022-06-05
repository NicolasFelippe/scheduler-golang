package user

import (
	"scheduler/src/db"
	"scheduler/src/models"
	"scheduler/src/repositories"
)

func GetById(id uint64) (models.User, error) {

	db, err := db.Connect()
	if err != nil {
		return models.User{}, err
	}

	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user, err := repository.GetById(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
