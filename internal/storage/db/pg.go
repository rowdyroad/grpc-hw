package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"math"
	"time"
)

type DB struct {
	db *pg.DB
}

const maxLimit = 100

func NewDB(config pg.Options) (storage.IStorage,error) {
	db := pg.Connect(&config)
	if err := db.Ping(context.TODO()); err != nil {
		return nil, err
	}
	ret := &DB{db}
	return ret, nil
}

func (c DB) GetTotalCount(from, to time.Time, low, high float64) (int,error) {
	query := c.db.Model((*Record)(nil))
	if !math.IsInf(low, -1) {
		query = query.Where("value >= ?", low)
	}
	if !math.IsInf(high, 1) {
		query = query.Where("value <= ?", high)
	}
	if !from.IsZero() {
		query = query.Where("time >= ?", from)
	}
	if !to.IsZero() {
		query = query.Where("time <= ?", to)
	}
	return query.Order("time asc").Count()
}

func (c DB) GetList(from ,to time.Time, low, high float64, offset,limit uint64) ([]storage.Record, error) {
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
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
	if err := query.Order("time asc").Offset(int(offset)).Limit(int(limit)).Select(); err != nil {
		return nil, err
	}
	ret := make([]storage.Record, len(records))
	for i, record := range records  {
		ret[i] = storage.Record{Time: record.Time, Value: record.Value}
	}
	return ret, nil
}

func (c DB) GetValue(time time.Time) (float64, error) {
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


func (c DB) GetDailyStats(from, to time.Time) (map[time.Time]storage.Stat, error) {
	items := []struct {
		storage.Stat
		Time time.Time
	}{}
	where := ""
	args := []interface{}{}
	if !from.IsZero() {
		where += " AND time >= ?"
		args = append(args, from)
	}
	if !to.IsZero() {
		where += " AND time <= ?"
		args = append(args, to)
	}
	_,err := c.db.Query(&items, `SELECT date_trunc('day', "time") "time", count(*) "count", avg(value) "average", min(value) "min", max(value) "max" FROM records WHERE 1=1 `+ where +` GROUP BY date_trunc('day', "time")`, args...)
	if err != nil {
		return nil, err
	}
	ret := map[time.Time]storage.Stat{}
	for _, item := range items {
		ret[item.Time] = item.Stat
	}
	return ret, nil
}