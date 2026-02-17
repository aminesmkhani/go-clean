package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP with time zone; not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; null"`
	DeletedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; null"`

	CreatedBy uint `gorm:"not null"`
	ModifiedBy int `gorm:"null"`
	DeletedBy int `gorm:"null"`
}