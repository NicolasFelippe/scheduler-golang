package scheduler

import (
	"scheduler/src/db"
	"scheduler/src/models"
	"scheduler/src/repositories"
)

func GetAll() ([]models.Scheduler, error) {

	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	repository := repositories.NewRepositoryScheduler(db)
	schedulers, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	return schedulers, nil
}
