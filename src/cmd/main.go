package main

import (
	"github.com/aminesmkhani/go-clean/api"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/cache"
	"github.com/aminesmkhani/go-clean/data/db"
	"github.com/aminesmkhani/go-clean/data/db/migrations"
	"github.com/aminesmkhani/go-clean/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization

func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(),nil)
	}

	err = db.InitDB(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(),nil)
	}

	migrations.Up1()

	api.InitServer(cfg)
}