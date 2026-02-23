package models



type Role struct {
	BaseModel
	Name string `gorm:"type:string;size:20;unique;not null"`
	Description string `gorm:"type:string;size:255;null"`
}