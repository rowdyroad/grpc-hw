package PG

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"math"
	"time"
)

type PG struct {
	db *pg.DB
}

func NewPG(config pg.Options) (storage.IStorage,error) {
	db := pg.Connect(&config)
	if err := db.Ping(context.TODO()); err != nil {
		return nil, err
	}
	return &PG{db}, nil
}

func (c PG) GetTotalCount(from, to time.Time, low, high float64) (int,error) {
	query := c.db.Model((*Record)(nil))
	if !math.IsInf(low, -1) {
		query = query.Where("value >= ?", low)
	}
	if !math.IsInf(low, 1) {
		query = query.Where("value <= ?", high)
	}
	if !from.IsZero() {
		query = query.Where("time >= ?", from)
	}
	if !from.IsZero() {
		query = query.Where("time <= ?", to)
	}
	return query.Count()
}

func (c PG) GetList(from ,to time.Time, low, high float64, offset,limit int) ([]storage.Record, error) {
	records := make([]Record, 0)
	query := c.db.Model(&records)
	if !math.IsInf(low, -1) {
		query = query.Where("value >= ?", low)
	}
	if !math.IsInf(low, 1) {
		query = query.Where("value <= ?", high)
	}
	if !from.IsZero() {
		query = query.Where("time >= ?", from)
	}
	if !from.IsZero() {
		query = query.Where("time <= ?", to)
	}
	if err := query.Offset(offset).Limit(limit).Select(); err != nil {
		return nil, err
	}
	ret := make([]storage.Record, len(records))
	for i, record := range records  {
		ret[i] = storage.Record{Time: record.Time, Value: record.Value}
	}
	return ret, nil
}

func (c PG) GetValue(time time.Time) (float64, error) {
	var record Record
	if err := c.db.Model(&record).Where("time = ?", time).Select(); err != nil {
		switch err {
		case pg.ErrNoRows:
			return 0, storage.ErrValueNotFound
		default:
			return 0, err
		}
	}
	return record.Value, nil
}


func (c PG) GetDailyStats(from, to time.Time) (map[time.Time]storage.Stat, error) {
	ret := map[time.Time]storage.Stat{}

	return ret, nil
}