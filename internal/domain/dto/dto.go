package dto

import "time"

type PricesDTO struct {
	Dates  []time.Time
	Names  []string
	Prices []float64
}
