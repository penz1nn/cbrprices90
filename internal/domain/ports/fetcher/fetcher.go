package fetcher

import (
	"cbrprices/internal/domain/dto"
	"time"
)

type Fetcher interface {
	Fetch(startDate time.Time, endDate time.Time) dto.PricesDTO
}
