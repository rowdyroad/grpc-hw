package storage

import (
	"errors"
	"time"
)

type Record struct {
	Time time.Time `json:"time"`
	Value float64 `json:"value"`
}

type Stat struct {
	Count uint64  `json:"count"`
	Average float64 `json:"average"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

var ErrValueNotFound = errors.New("value not found")

type IStorage interface {
	GetTotalCount(from,to time.Time, low, high float64) (int,error)
	GetList(from ,to time.Time, low, high float64, offset,limit uint64) ([]Record,error)
	GetValue(time time.Time) (float64,error)
	GetDailyStats(from, to time.Time) (map[time.Time]Stat, error)
}
