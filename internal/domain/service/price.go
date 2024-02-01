package service

import (
	"local/cbrprices/internal/domain/model"
	"local/cbrprices/internal/domain/ports/fetcher"
	"time"
)

type Service struct {
	f fetcher.Fetcher
}

func (s Service) GetResults(startDate time.Time, endDate time.Time) []model.CurrencyResult {
	data := s.f.Fetch(startDate, endDate)

	results := []model.CurrencyResult{}
	for currency, prices := range data.Prices {
		result := model.CurrencyResult{
			Name: currency,
		}

		highPriceIndex := 0
		lowPriceIndex := 0
		var sum float64 = 0
		for index, price := range prices {
			sum = sum + price
			if price > prices[highPriceIndex] {
				highPriceIndex = index
			}
			if price < prices[lowPriceIndex] {
				lowPriceIndex = index
			}
		}
		result.DateHighPrice = data.Dates[highPriceIndex]
		result.HighPrice = prices[highPriceIndex]
		result.DateLowPrice = data.Dates[lowPriceIndex]
		result.LowPrice = prices[lowPriceIndex]
		result.AveragePrice = sum / float64(len(prices))
		results = append(results, result)
	}
	return results
}
