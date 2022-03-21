package schema

import "time"

type ListRequest struct {
	From time.Time `form:"from" example:"2011-05-03T11:58:01+03:00"`
	To time.Time `form:"to" example:"2011-05-03T11:58:01+03:00"`
	Low *float64 `form:"low" example:"105.1"`
	High *float64 `form:"high" example:"201"`
	Offset int `form:"offset" example:"0"`
	Limit int `form:"limit" example:"20"`
}

type ValueRequest struct {
	Time time.Time `uri:"time" example:"2011-05-03T11:58:01+03:00"`
}

type StatsRequest struct {
	From time.Time `form:"from" example:"2011-05-03T11:58:01+03:00"`
	To time.Time `form:"to" example:"2011-05-03T11:58:01+03:00"`
}
