package database

import (
	"github.com/aminesmkhani/go-clean/config"
	"gorm.io/gorm"
)


var DbClient *gorm.DB



func InitDB(cfg config.Config) error{
	
}