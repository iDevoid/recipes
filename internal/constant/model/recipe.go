package model

import "time"

type Recipe struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	MakingTime  string    `json:"making_time" db:"making_time"`
	Serves      string    `json:"serves" db:"serves"`
	Ingredients string    `json:"ingredients" db:"ingredients"`
	Cost        int32     `json:"cost" db:"cost"`
	CreateAt    time.Time `json:"created_at" db:"created_at"`
	UpdateAt    time.Time `json:"updated_at" db:"updated_at"`
}

func (r *Recipe) Valid() bool {
	if r.Title == "" || r.MakingTime == "" || r.Serves == "" || r.Ingredients == "" || r.Cost < 1 {
		return false
	}
	return true
}

type SuccessResp struct {
	Message string   `json:"message"`
	Recipes []Recipe `json:"recipe"`
}

type FailedResp struct {
	Message  string `json:"message"`
	Required string `json:"required"`
}
