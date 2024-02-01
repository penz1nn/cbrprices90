package fetcher

import (
	"local/cbrprices/pkg/logger"
	"log/slog"
	"testing"
	"time"
)

func TestFetcher_New(t *testing.T) {
	log := logger.SetupLogger("local")
	f := New(log)
	if f.client == nil {
		t.Error("http client is nil")
	}
}

func TestFetcher_getDayData(t *testing.T) {
	log := logger.SetupLogger("local")
	f := New(log)
	time, err := time.Parse("04.01.2006", "15.01.2024")
	if err != nil {
		t.Error(err)
	}
	data := f.getDayData(time)
	log.Debug("got data", slog.Any("data", data))
}

func TestFetcher_Fetch(t *testing.T) {
	log := logger.SetupLogger("local")
	f := New(log)
	time2 := time.Now().Add(-24 * time.Hour)
	time1 := time2.Add(-30 * 24 * time.Hour)
	if time1.After(time2) {
		t.Errorf("Wrong times: %v, %v", time1, time2)
	}

	data := f.Fetch(time1, time2)
	if len(data.Dates) < 1 || len(data.Prices) < 1 {
		t.Error("Got empty data")
	}
	if len(data.Dates) != 30 {
		t.Errorf("Should have got 30 days of data, got %d", len(data.Dates))
	}
	for currency, prices := range data.Prices {
		if len(prices) != 30 {
			t.Errorf("Should have 30 days of data but %s got %d", currency, len(prices))
		}
	}
}
