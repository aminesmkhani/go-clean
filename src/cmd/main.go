package main

import (
	"github.com/aminesmkhani/go-clean/api"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/cache"
)

func main(){
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}