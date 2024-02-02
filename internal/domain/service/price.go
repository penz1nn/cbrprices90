package service

import (
	"cbrprices/internal/domain/model"
	"cbrprices/internal/domain/ports/fetcher"
	"math"
	"time"
)

type Service struct {
	f fetcher.Fetcher
}

type allPrices struct {
	Dates  []time.Time
	Prices []float64
}

func (s Service) GetResult(startDate time.Time, endDate time.Time) (result model.Result) {
	data := s.f.Fetch(startDate, endDate)
	result.MaxRatePrice = 0
	result.MinRatePrice = math.MaxFloat64
	var rubPriceSum float64 = 0

	for index, price := range data.Prices {
		if price > result.MaxRatePrice {
			result.MaxRatePrice = price
			result.MaxRateCurrency = data.Names[index]
			result.MaxRateDate = data.Dates[index]
		}
		if price < result.MinRatePrice {
			result.MinRatePrice = price
			result.MinRateCurrency = data.Names[index]
			result.MinRateDate = data.Dates[index]
		}
		rubPriceSum = rubPriceSum + 1/price
	}
	result.AverageRubRate = rubPriceSum / float64(len(data.Prices))
	return
}

func NewService(f fetcher.Fetcher) Service {
	return Service{
		f: f,
	}
}
