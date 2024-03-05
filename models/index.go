package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID        uint
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
