package csv

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestNewCSV(t *testing.T) {
	csv,err := NewCSV("../../../data/meterusage.csv")
	if err != nil {
		t.Fatal(err)
	}
	count, err := csv.GetTotalCount(
		time.Time{},
		time.Time{},
		math.Inf(-1),
		math.Inf(1),
	)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, 2974, count) {
		return
	}

	count, err = csv.GetTotalCount(
		time.Date(2019,time.January, 1, 20,45,0,0, time.UTC),
		time.Date(2019,time.January, 2, 21,45,0, 0, time.UTC),
		math.Inf(-1),
		math.Inf(1),
	)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, 101, count) {
		return
	}

	count, err = csv.GetTotalCount(
		time.Date(2019,time.January, 1, 20,45,0,0, time.UTC),
		time.Date(2019,time.January, 2, 21,45,0, 0, time.UTC),
		190,
		math.Inf(1),
	)
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Equal(t, 41, count) {
		return
	}
	for i := 0; i < count;i+=10 {
		data,err := csv.GetList(
			time.Date(2019, time.January, 1, 20, 45, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 21, 45, 0, 0, time.UTC),
			200,
			math.Inf(1),
			i, 10,
		)
		if !assert.NoError(t, err) {
			return
		}
		if !assert.True(t, len(data) > 0 && len(data) <= 10) {
			return
		}

		for _,record := range data {
			if !assert.True(t, record.Value >= 200) {
				return
			}

			value,err := csv.GetValue(record.Time)
			if !assert.NoError(t, err) {
				return
			}
			if !assert.Equal(t, record.Value, value) {
				return
			}
		}
	}

	_, err = csv.GetValue(time.Now())

	assert.Error(t, err)

	stats, err := csv.GetDailyStats(
		time.Date(2019, time.January, 2, 16,00,0,0, time.UTC),
		time.Date(2019, time.January, 2, 18,00,0, 0, time.UTC),
	)
	day := time.Date(2019,time.January, 2, 0,00,0,0, time.UTC)

	if !assert.Equal(t, 9, stats[day].Count) {
		return
	}
	if !assert.Equal(t, 155.34, stats[day].Min) {
		return
	}
	if !assert.Equal(t, 240.96, stats[day].Max) {
		return
	}
	if !assert.Equal(t, 215.047, math.Round(stats[day].Average * 1000) / 1000) {
		return
	}
}

