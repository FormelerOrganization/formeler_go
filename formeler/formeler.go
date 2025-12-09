package formeler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const DEFAULT_BASE_URL = "https://formeler.com/api"
const API_VERSION = "2025-11-12"

type Formeler struct {
	LangID LangID
}

type LangID struct {
	apiKey  string
	baseURL string
}

func NewFormeler(apiKey string) *Formeler {
	f := Formeler{LangID: LangID{apiKey: apiKey, baseURL: DEFAULT_BASE_URL}}
	return &f
}

func NewFormelerWithURL(apiKey string, baseURL string) *Formeler {
	f := Formeler{LangID: LangID{apiKey: apiKey, baseURL: baseURL}}
	return &f
}

func (l LangID) Detect(text string) (string, error) {
	data, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", l.baseURL+"/"+API_VERSION+"/langid/detect", strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", l.apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (l LangID) BatchDetect(texts []string) (string, error) {
	data, err := json.Marshal(map[string][]string{"texts": texts})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", l.baseURL+"/"+API_VERSION+"/langid/batch-detect", strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", l.apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
