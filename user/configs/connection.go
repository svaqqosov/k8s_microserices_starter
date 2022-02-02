package config

import (
	"os"

	"github.com/sirupsen/logrus"
	model "github.com/svaqqosov/k8s_microserices_starter/models"
	util "github.com/svaqqosov/k8s_microserices_starter/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("SVC_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("SVC_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
