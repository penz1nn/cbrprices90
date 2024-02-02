package main

import (
	"bytes"
	"cbrprices/internal/domain/model"
	"encoding/json"
	"fmt"
)

func writeResults(result model.Result) {
	resultsJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, resultsJson, "", "	"); err != nil {
		panic(err)
	}
	fmt.Println(prettyJson.String())
}
