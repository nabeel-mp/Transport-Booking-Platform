package models

import (
	"time"

	"github.com/google/uuid"
)

// TrainSchedule is one actual dated run of a Train.
type TrainSchedule struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TrainID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Train        Train     `gorm:"foreignKey:TrainID"`
	ScheduleDate time.Time `gorm:"type:date;not null"`
	DepartureAt  time.Time `gorm:"not null"` // full datetime with timezone
	ArrivalAt    time.Time `gorm:"not null"`
	Status       string    `gorm:"size:20;not null;default:'SCHEDULED'"` // SCHEDULED|DEPARTED|IN_TRANSIT|ARRIVED|CANCELLED|DELAYED
	DelayMinutes int       `gorm:"default:0"`
	AvailableSL  int       `gorm:"not null;default:0"`
	Available3AC int       `gorm:"not null;default:0"`
	Available2AC int       `gorm:"not null;default:0"`
	Available1AC int       `gorm:"not null;default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (TrainSchedule) TableName() string {
	return "train_schedules"
}
