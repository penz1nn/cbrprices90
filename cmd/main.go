package main

import (
	"cbrprices/internal/adapters/fetcher"
	"cbrprices/internal/domain/service"
	"cbrprices/pkg/logger"
	"flag"
	"time"
)

func main() {
	dateStartStr := flag.String(
		"start",
		time.Now().Add(-90*24*time.Hour).Format("02-01-2006"),
		"Start time in format DD-MM-YYYY",
	)
	dateEndStr := flag.String(
		"end",
		time.Now().Format("02-01-2006"),
		"End time in format DD-MM-YYYY",
	)
	flag.Parse()

	startDate, err := time.Parse("02-01-2006", *dateStartStr)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("02-01-2006", *dateEndStr)
	if err != nil {
		panic(err)
	}

	log := logger.SetupLogger("prod")
	fetcher := fetcher.New(log)
	service := service.NewService(fetcher)
	data := service.GetResult(startDate, endDate)

	writeResults(data)
}
