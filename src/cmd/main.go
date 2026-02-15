package main

import (
	"log"

	"github.com/aminesmkhani/go-clean/api"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/cache"
	"github.com/aminesmkhani/go-clean/data/db"
	"github.com/aminesmkhani/go-clean/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization

func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDB(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}