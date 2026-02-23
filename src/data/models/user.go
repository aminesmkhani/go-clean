package models


type User struct {
	BaseModel
	Username 		string  `gorm:"type:string;size:20;unique;not null"`
	FirstName 		string	`gorm:"type:string;size:15;null"`
	LastName 		string	`gorm:"type:string;size:25;null"`
	MobileNUmber	string 	`gorm:"type:string;size:11;null;unique;default:null"`
	Email 			string	`gorm:"type:string;size:64;null;unique;default:null"`
	Password 		string	`gorm:"type:string;size:64;not null"`
	IsActive 		bool	`gorm:"type:boolean;default:true"`
	UserRoles 		*[]UserRole
}