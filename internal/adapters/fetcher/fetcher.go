package fetcher

import (
	"bytes"
	"cbrprices/internal/domain/dto"
	"encoding/xml"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

type fetcher struct {
	logger *slog.Logger
	client *http.Client
}

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"ID,attr"`
		NumCode   string `xml:"NumCode"`   // 036, 944, 826, 051, 933, ...
		CharCode  string `xml:"CharCode"`  // AUD, AZN, GBP, AMD, BYN, ...
		Nominal   string `xml:"Nominal"`   // 1, 1, 1, 100, 1, 1, 1, 10...
		Name      string `xml:"Name"`      // Australian Dollar, Azerba...
		Value     string `xml:"Value"`     // 55,6634, 44,9663, 100,959...
		VunitRate string `xml:"VunitRate"` // 55,6634, 44,9663, 100,959...
	} `xml:"Valute"`
}

func (f fetcher) Fetch(startDate time.Time, endDate time.Time) dto.PricesDTO {
	op := "fetcher/Fetch"
	f.logger.With("operation", op)
	result := dto.PricesDTO{}
	result.Prices = map[string][]float64{}

	date := startDate
	for date.Before(endDate) {
		result.Dates = append(result.Dates, date)
		f.logger.Debug(
			"getting data for one day...",
			slog.Any("date", date.Format("2006.01.15")),
		)
		data := f.getDayData(date)
		for _, valute := range data.Valute {
			price, err := strconv.ParseFloat(
				strings.ReplaceAll(valute.VunitRate, ",", "."),
				64,
			)
			if err != nil {
				f.logger.Error("error reading float64 from XML data", slog.Any("error", err))
				return dto.PricesDTO{}
			}
			f.logger.Debug("got price", slog.String("currency", valute.CharCode), slog.Float64("price", price))
			result.Prices[valute.CharCode] = append(result.Prices[valute.CharCode], price)
		}
		date = date.Add(time.Hour * 24)
	}
	f.logger.Debug("got data", slog.Any("data", result))
	return result
}

func (f fetcher) getDayData(date time.Time) valCurs {
	op := "fetcher/getDayData"
	f.logger.With("operation", op)
	urlBase := "http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req="
	dateStr := date.Format("02/01/2006")
	url := urlBase + dateStr

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0")
	if err != nil {
		f.logger.Error("could not make request", slog.String("url", url), slog.Any("error", err))
		return valCurs{}
	}

	response, err := f.client.Do(request)
	if err != nil {
		f.logger.Error("error making a request", slog.String("url", url), slog.Any("error", err))
		return valCurs{}
	}
	f.logger.Debug("got HTTP response", slog.String("url", url), slog.Int("status", response.StatusCode))

	body, err := io.ReadAll(response.Body)
	if err != nil {
		f.logger.Error("error reading HTTP response", slog.Any("error", err))
		return valCurs{}
	}

	var result valCurs
	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Decode(&result)
	if err != nil {
		f.logger.Error("error unmarshalling XML", slog.Any("error", err))
		return valCurs{}
	}
	return result
}

func New(log *slog.Logger) fetcher {
	return fetcher{
		logger: log,
		client: &http.Client{},
	}
}
