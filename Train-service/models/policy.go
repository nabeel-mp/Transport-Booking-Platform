package models

import (
	"time"

	"github.com/google/uuid"
)

// CancellationPolicy defines refund rules based on hours before departure.
type CancellationPolicy struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name                 string    `gorm:"size:100;not null"`
	HoursBeforeDeparture int       `gorm:"not null"`                   // applies if cancelled >= X hours before
	RefundPercentage     float64   `gorm:"type:decimal(5,2);not null"` // 90.00 | 50.00 | 0.00
	CancellationFee      float64   `gorm:"type:decimal(10,2);not null;default:0"`
	IsActive             bool      `gorm:"default:true"`
	CreatedAt            time.Time
}

// Cancellation is created when a user requests cancellation of a booking.
type Cancellation struct {
	ID              uuid.UUID           `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID       uuid.UUID           `gorm:"type:uuid;uniqueIndex;not null"`
	Booking         TrainBooking        `gorm:"foreignKey:BookingID"`
	Reason          string              `gorm:"type:text"`
	RefundAmount    float64             `gorm:"type:decimal(10,2);not null"`
	RefundStatus    string              `gorm:"size:20;not null;default:'PENDING'"` // PENDING|PROCESSING|COMPLETED|FAILED
	PolicyAppliedID *uuid.UUID          `gorm:"type:uuid"`
	PolicyApplied   *CancellationPolicy `gorm:"foreignKey:PolicyAppliedID"`
	RequestedAt     time.Time           `gorm:"default:now()"`
	ProcessedAt     *time.Time
	CreatedAt       time.Time
}

func (CancellationPolicy) TableName() string { return "cancellation_policies" }
func (Cancellation) TableName() string       { return "cancellations" }
