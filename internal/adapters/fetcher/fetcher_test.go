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
	log.Info("got data", slog.Any("data", data))
}
