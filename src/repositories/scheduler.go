package repositories

import (
	"database/sql"
	"scheduler/src/models"
)

type schedulers struct {
	db *sql.DB
}

func NewRepositoryScheduler(db *sql.DB) *schedulers {
	return &schedulers{db}
}

func (repository schedulers) GetAll() ([]models.Scheduler, error) {
	rows, err := repository.db.Query("select sc.id, sc.user_id, sc.client_id, sc.date from schedulers sc where sc.date > now() and Date(sc.date) = curdate();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedulers []models.Scheduler

	for rows.Next() {
		var scheduler models.Scheduler
		if err = rows.Scan(
			&scheduler.ID,
			&scheduler.UserId,
			&scheduler.ClientId,
			&scheduler.Date,
		); err != nil {
			return nil, err
		}

		schedulers = append(schedulers, scheduler)
	}

	return schedulers, nil
}
