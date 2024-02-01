package dto

import "time"

type PricesDTO struct {
	Dates  []time.Time
	Prices map[string][]float64
}
