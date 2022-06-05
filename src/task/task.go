package task

import (
	"fmt"
	"scheduler/src/services/client"
	"scheduler/src/services/company"
	"scheduler/src/services/scheduler"
	"scheduler/src/services/user"
	"scheduler/src/sqs"
	"time"
)

const (
	DDMMYYYYhhmm = "02/01/2006 15:04"
)

func SendNotificationSchedulers() {
	fmt.Println("I am running task.")
	schedulers, err := scheduler.GetAll()
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime)
	fmt.Println("schedulers", schedulers)

	for _, scheduler := range schedulers {
		diff := scheduler.Date.Sub(currentTime).Minutes()
		fmt.Println("diff", diff)
		if diff <= 250 && diff >= 50 {
			client, err := client.GetById(scheduler.ClientId)
			if err != nil {
				panic(err)
			}
			fmt.Println("client", client)
			user, err := user.GetById(scheduler.UserId)
			if err != nil {
				panic(err)
			}
			fmt.Println("client", user)
			company, err := company.GetById(user.CompanyId)
			if err != nil {
				panic(err)
			}
			fmt.Println("company", company)
			bodyMessage := fmt.Sprintf("Olá %s, você tem um encontro marcado em %s às %s", client.Name, company.Name, scheduler.Date.Format(DDMMYYYYhhmm))
			sqs.Send(company.Name, client.Phone, bodyMessage)
		}
	}
}
