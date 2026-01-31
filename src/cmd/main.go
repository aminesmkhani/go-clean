package main

import (
	"log"

	"github.com/aminesmkhani/go-clean/api"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/cache"
	"github.com/aminesmkhani/go-clean/data/db"
)

func main(){
	cfg := config.GetConfig()

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