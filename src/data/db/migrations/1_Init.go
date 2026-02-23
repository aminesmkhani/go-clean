package migrations

import (
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/db"
	"github.com/aminesmkhani/go-clean/data/models"
	"github.com/aminesmkhani/go-clean/pkg/logging"
)


var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()

	database.AutoMigrate(&models.Country{})
	database.AutoMigrate(&models.City{})

	logger.Info(logging.Postgres, logging.Migration, "Migration 1 completed successfully",nil) 

}

func Down1() {

}
