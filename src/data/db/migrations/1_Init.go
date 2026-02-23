package migrations

import (
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/constants"
	"github.com/aminesmkhani/go-clean/data/db"
	"github.com/aminesmkhani/go-clean/data/models"
	"github.com/aminesmkhani/go-clean/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()

	createTable(database) 
	createDefaultInformation(database)

}

func createTable(database *gorm.DB) {

	database.AutoMigrate(&models.Country{})
	database.AutoMigrate(&models.City{})
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Role{})
	database.AutoMigrate(&models.UserRole{})

	logger.Info(logging.Postgres, logging.Migration, "Migration 1 completed successfully", nil)
}



func createDefaultInformation(database *gorm.DB) {


	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database,&adminRole)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database,&defaultRole)

	u := models.User{Username: constants.DefaultUserName, 
		FirstName: "Amin",
		LastName: "Esmkhani",
		MobileNUmber: "09123456787",
		Email: "amin.esmkhani75@gmail.com",
	}
	pass := "12345678"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	createUserIfNotExists(database,&u,adminRole.Id)

	logger.Info(logging.Postgres, logging.Migration, "Default information created successfully", nil)
}

func createRoleIfNotExists(database *gorm.DB,r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createUserIfNotExists(database *gorm.DB,u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserID: u.Id, RoleID: roleId}
		database.Create(&ur)
	}
}

func Down1() {

}
