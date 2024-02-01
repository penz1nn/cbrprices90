package model

import "time"

type CurrencyResult struct {
	Name          string    `json:"name"`
	HighPrice     float32   `json:"highPrice"`
	DateHighPrice time.Time `json:"dateHighPrice"`
	LowPrice      float32   `json:"lowPrice"`
	DateLowPrice  time.Time `json:"dateLowPrice"`
	AveragePrice  float32   `json:"averagePrice"`
}
