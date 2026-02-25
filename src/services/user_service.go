package services

import (
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/data/db"
	"github.com/aminesmkhani/go-clean/pkg/logging"
	"gorm.io/gorm"
)



type UserService struct {
	logger logging.Logger
	cfg 	*config.Config
	otpService *OtpService
	database   *gorm.DB
}



func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		logger: logger,
		cfg: cfg,
		database: database,
		otpService: NewOtpService(cfg),
	}
}


func (s *UserService) SendOtp()