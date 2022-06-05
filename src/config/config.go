package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConnectionDb = ""
	RegionAws          = ""
	AccessKeyId        = ""
	SecretAccessKey    = ""
	QueueUrl           = ""
)

func LoadConfiguration() {
	var er error
	if er = godotenv.Load(); er != nil {
		log.Fatal(er)
	}
	StringConnectionDb = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
	)

	RegionAws = os.Getenv("REGION_AWS")
	AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	QueueUrl = os.Getenv("QUEUE_URL")

}
