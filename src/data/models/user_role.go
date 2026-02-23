package models


type UserRole struct {
	BaseModel
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	UserID int `gorm:"not null"`
	RoleID int `gorm:"not null"`
}