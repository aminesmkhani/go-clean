package models


type City struct {
	BaseModel
	Name string `gorm:"type:string;size:10;not null;" json:"name"`
	CountryId int `gorm:"not null" json:"countryId"`
	Country Country `gorm:"foreignKey:CountryId" json:"country"`
}