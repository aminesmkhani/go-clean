package models

type Country struct {
	BaseModel
	Name string `gorm:"type:string;size:15;not null;" json:"name"`
	Code string `gorm:"type:string;size:2;not null" json:"code"`
}