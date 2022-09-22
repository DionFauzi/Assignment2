package models

import (
	"time"
)

type Brand struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Cars      []Car
}
