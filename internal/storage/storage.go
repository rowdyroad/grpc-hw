package storage

import (
	"errors"
	"time"
)

type Record struct {
	Time time.Time
	Value float64
}

type Stat struct {
	Count int
	Average float64
	Min float64
	Max float64
}

var ErrValueNotFound = errors.New("value not found")

type IStorage interface {
	GetTotalCount(from,to time.Time, low, high float64) (int,error)
	GetList(from ,to time.Time, low, high float64, offset,limit int) ([]Record,error)
	GetValue(time time.Time) (float64,error)
	GetDailyStats(time, to time.Time) (map[time.Time]Stat, error)
}
