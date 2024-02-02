package main

import (
	"bytes"
	"cbrprices/internal/domain/model"
	"encoding/json"
	"fmt"
)

func writeResults(results []model.CurrencyResult) {
	resultsJson, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, resultsJson, "", "	"); err != nil {
		panic(err)
	}
	fmt.Println(prettyJson.String())
}
