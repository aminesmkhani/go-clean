package services

import (
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/cache"
	"github.com/aminesmkhani/go-clean/pkg/logging"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}


func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{
		logger: logger,
		cfg:    cfg,
		redis:  redis,
	}
}
