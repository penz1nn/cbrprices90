package model

import "time"

type CurrencyResult struct {
	Name          string    `json:"name"`
	HighPrice     float64   `json:"highPrice"`
	DateHighPrice time.Time `json:"dateHighPrice"`
	LowPrice      float64   `json:"lowPrice"`
	DateLowPrice  time.Time `json:"dateLowPrice"`
	AveragePrice  float64   `json:"averagePrice"`
}
