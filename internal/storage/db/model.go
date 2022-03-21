package db

import "time"

type Record struct {
	tableName struct{} `pg:"records, discard_unknown_columns"`

	Time time.Time `pg:"time"`
	Value float64 `pg:"value"`
}
