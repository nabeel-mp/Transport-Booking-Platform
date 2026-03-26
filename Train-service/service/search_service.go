package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nabeel-mp/tripneo/train-service/db"
	"github.com/nabeel-mp/tripneo/train-service/models"
	"github.com/nabeel-mp/tripneo/train-service/repository"
	goredis "github.com/redis/go-redis/v9"
)

const searchCacheTTL = 2 * time.Minute

func SearchTrains(ctx context.Context, rdb *goredis.Client, fromCode, toCode, date, class string) ([]models.TrainSchedule, error) {
	var results []models.TrainSchedule

	// Use the 'class' parameter in your query if needed
	query := db.DB.WithContext(ctx).Table("train_schedules").
		Select("train_schedules.*").
		Joins("JOIN trains ON trains.id = train_schedules.train_id").
		Joins("JOIN train_stops AS s1 ON s1.train_id = trains.id").
		Joins("JOIN stations AS st1 ON st1.id = s1.station_id").
		Joins("JOIN train_stops AS s2 ON s2.train_id = trains.id").
		Joins("JOIN stations AS st2 ON st2.id = s2.station_id").
		Where("st1.code = ? AND st2.code = ?", fromCode, toCode).
		Where("s1.stop_sequence < s2.stop_sequence").
		Where("DATE(train_schedules.schedule_date) = ?", date)

	// Optional: Filter by class if provided
	if class != "" {
		// Assuming your model/DB has a way to filter by class
		query = query.Where("class = ?", class)
	}

	err := query.Preload("Train.Stops.Station").Find(&results).Error

	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return results, nil
}

// GetScheduleDetail returns a single schedule with its train details.
func GetScheduleDetail(scheduleID string) (interface{}, error) {
	schedule, err := repository.GetScheduleByID(scheduleID)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func GetSeatMap(
	ctx context.Context,
	rdb *goredis.Client,
	scheduleID, class string,
) (interface{}, error) {
	seats, err := repository.GetSeatsByScheduleAndClass(scheduleID, class)
	if err != nil {
		return nil, err
	}

	type SeatWithLock struct {
		ID         string  `json:"id"`
		SeatNumber string  `json:"seat_number"`
		Coach      string  `json:"coach"`
		Class      string  `json:"class"`
		BerthType  string  `json:"berth_type"`
		Status     string  `json:"status"`
		Price      float64 `json:"price"`
		IsLocked   bool    `json:"is_locked"`
	}

	result := make([]SeatWithLock, len(seats))
	for i, s := range seats {
		isLocked := false
		if s.Status == "AVAILABLE" {
			// Check Redis lock — display only, not for booking gate
			locked, _ := checkLockStatus(ctx, rdb, scheduleID, s.ID.String())
			isLocked = locked
		}
		result[i] = SeatWithLock{
			ID:         s.ID.String(),
			SeatNumber: s.SeatNumber,
			Coach:      s.Coach,
			Class:      s.Class,
			BerthType:  s.BerthType,
			Status:     s.Status,
			Price:      s.Price,
			IsLocked:   isLocked,
		}
	}
	return result, nil
}

// checkLockStatus is an internal helper that reads the Redis lock key.
func checkLockStatus(ctx context.Context, rdb *goredis.Client, scheduleID, seatID string) (bool, error) {
	key := fmt.Sprintf("seat:lock:train:%s:%s", scheduleID, seatID)
	_, err := rdb.Get(ctx, key).Result()
	if err == goredis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
