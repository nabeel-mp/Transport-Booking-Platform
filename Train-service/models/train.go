package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Train struct {
	ID                 uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TrainNumber        string        `gorm:"size:10;uniqueIndex;not null"` // 12678
	TrainName          string        `gorm:"size:200;not null"`            // Ernakulam Express
	OriginStation      string        `gorm:"size:5;not null"`              // ERS
	DestinationStation string        `gorm:"size:5;not null"`              // MAS
	DepartureTime      string        `gorm:"size:5;not null"`              // "20:30" stored as HH:MM
	ArrivalTime        string        `gorm:"size:5;not null"`              // "05:30"
	DurationMinutes    int           `gorm:"not null"`
	DaysOfWeek         pq.Int32Array `gorm:"type:integer[];not null"` // {1,2,3,4,5,6,7}
	IsActive           bool          `gorm:"default:true"`
	CreatedAt          time.Time
}
