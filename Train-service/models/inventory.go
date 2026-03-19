package models

import (
	"time"

	"github.com/google/uuid"
)

// TrainInventory represents one pre-purchased confirmed berth
// for a specific TrainSchedule.
type TrainInventory struct {
	ID              uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TrainScheduleID uuid.UUID     `gorm:"type:uuid;not null;index"`
	TrainSchedule   TrainSchedule `gorm:"foreignKey:TrainScheduleID"`
	SeatNumber      string        `gorm:"size:10;not null"`                     // 23
	Coach           string        `gorm:"size:10;not null"`                     // S1
	Class           string        `gorm:"size:5;not null"`                      // SL | 3AC | 2AC | 1AC
	BerthType       string        `gorm:"size:15;not null"`                     // LOWER|MIDDLE|UPPER|SIDE_LOWER|SIDE_UPPER
	Status          string        `gorm:"size:20;not null;default:'AVAILABLE'"` // AVAILABLE|BOOKED|BLOCKED|LOCKED
	Price           float64       `gorm:"type:decimal(10,2);not null"`          // current selling price
	WholesalePrice  float64       `gorm:"type:decimal(10,2);not null"`          // what Tripneo paid
	ProviderID      string        `gorm:"size:100"`                             // provider booking reference
	QuantitySold    int           `gorm:"not null;default:0"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (TrainInventory) TableName() string {
	return "train_inventory"
}
