package main

import (
	"scheduler/src/config"
	"scheduler/src/task"

	"github.com/jasonlvhit/gocron"
)

func main() {
	config.LoadConfiguration()
	s := gocron.NewScheduler()
	s.Every(2).Minute().Do(task.SendNotificationSchedulers)
	<-s.Start()
}
