package cashback

import "time"

type Cashback struct {
	Percentage float64    `json:"percentage" binding:"required"`
	Limit      *int32     `json:"limit,omitempty"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
}
