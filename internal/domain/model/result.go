package model

import "time"

type Result struct {
	MaxRateCurrency string    `json:"maxRateCurrency"`
	MaxRatePrice    float64   `json:"maxRatePrice"`
	MaxRateDate     time.Time `json:"maxRateTime"`
	MinRateCurrency string    `json:"minRateCurrency"`
	MinRatePrice    float64   `json:"minRatePrice"`
	MinRateDate     time.Time `json:"minRateDate"`
	AverageRubRate  float64   `json:"averageRubRate"`
}
