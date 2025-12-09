package main

import (
	"encoding/json"
	"fmt"
	"formeler"
)

type ResultResponse struct {
	Result     string   `json:"result"`
	Message    string   `json:"message"`
	Language   string   `json:"language"`
	Languages  []string `json:"languages"`
	TokenCount int64    `json:"token_count"`
}

func detectExample() {
	formeler := formeler.NewFormeler("YOUR-API-KEY")
	result, err := formeler.LangID.Detect("Dies ist ein Test")
	if err != nil {
		panic(err)
	}
	var resultJSON ResultResponse
	err = json.Unmarshal([]byte(result), &resultJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultJSON.Result, resultJSON.Language, resultJSON.TokenCount)
}

func batchDetectExample() {
	formeler := formeler.NewFormeler("YOUR-API-KEY")
	result, err := formeler.LangID.BatchDetect([]string{"Dies ist ein Test", "this is a test"})
	if err != nil {
		panic(err)
	}
	var resultJSON ResultResponse
	err = json.Unmarshal([]byte(result), &resultJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultJSON.Result, resultJSON.Languages, resultJSON.TokenCount)
}

func main() {
	detectExample()
	batchDetectExample()
}
