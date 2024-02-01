package fetcher

import (
	"local/cbrprices/internal/dto"
	"time"
)

type Fetcher interface {
	Fetch(startDate time.Time, endDate time.Time) dto.PricesDTO
}
