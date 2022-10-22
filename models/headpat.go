package models

import "time"

type Headpat struct {
	ID   uint      `gorm:"primaryKey;not null"`
	Time time.Time `gorm:"not null"`
}
