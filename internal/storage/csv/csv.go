package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

const maxLimit = 100

type CSV struct {
	data []storage.Record
}

func NewCSV(filename string) (storage.IStorage,error) {
	file,err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]storage.Record,0)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		t, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			fmt.Println("Error parse time from row:", record)
			continue
		}
		v, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println("Error parse time from row:", record)
			continue
		}
		data = append(data, storage.Record{Time: t, Value: v})
	}
	sort.Slice(data, func(i,j int) bool {
		return data[i].Time.Before(data[j].Time)
	})
	return &CSV{data}, nil
}

func (c CSV) GetTotalCount(from, to time.Time, low, high float64) (int,error) {
	var count int
	ret := []storage.Record{}
	for _, record := range c.data {
		if (from.IsZero() || record.Time.Equal(from) || record.Time.After(from)) && (to.IsZero() || record.Time.Equal(to) || record.Time.Before(to)) {
			if record.Value >= low && record.Value <= high {
				count++
				ret = append(ret, record)
			}
		}
	}
	return count,nil
}

func (c CSV) GetList(from ,to time.Time, low, high float64, offset,limit uint64) ([]storage.Record,error) {
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	ret := make([]storage.Record, 0, limit)
	for _, record := range c.data {
		if (from.IsZero() || record.Time.Equal(from) || record.Time.After(from)) && (to.IsZero() || record.Time.Equal(to) || record.Time.Before(to)) {
			if record.Value >= low && record.Value <= high {
				ret = append(ret, record)
				if len(ret) == cap(ret) {
					break
				}
			}
		}
	}
	return ret,nil
}

func (c CSV) GetValue(time time.Time) (float64,error) {
	for _, record := range c.data {
		if record.Time.Equal(time) {
			return record.Value, nil
		}
	}
	return 0, storage.ErrValueNotFound
}

func (c CSV) GetDailyStats(from, to time.Time) (map[time.Time]storage.Stat, error) {
	ret := map[time.Time]storage.Stat{}
	for _, record := range c.data {
		if (from.IsZero() || record.Time.Equal(from) || record.Time.After(from)) && (to.IsZero() || record.Time.Equal(to) || record.Time.Before(to)) {
			y,m, d := record.Time.Date()
			date := time.Date(y,m,d,0,0,0,0,time.UTC)
			if stat, ok := ret[date]; !ok {
				ret[date] = storage.Stat{
					Count: 1,
					Average: record.Value,
					Min: record.Value,
					Max: record.Value,
				}
			} else {
				stat.Count++
				stat.Average += record.Value
				stat.Min = math.Min(stat.Min, record.Value)
				stat.Max = math.Max(stat.Max, record.Value)
				ret[date] = stat
			}
		}
	}
	for t, r := range ret {
		r.Average /= float64(r.Count)
		ret[t] = r
	}
	return ret, nil
}