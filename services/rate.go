package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	key    = "886314c0-c94d-41b9-a05c-9e2dcf1892d4"
)

type BTC struct {
	Data []struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Quote  struct {
			UAH struct {
				Price float64 `json:"price"`
			} `json:"UAH"`
		} `json:"quote"`
	} `json:"data"`
}

func GetBitcoinRate() (float64, error) {
	client := &http.Client{}
	req, err := createRequest()
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending request to server: %w", err)
	}
	defer resp.Body.Close()

	body, err := readResponseBody(resp)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	bit, err := unmarshalResponse(body)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return bit.Data[0].Quote.UAH.Price, nil
}

func createRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("start", "1")
	q.Add("limit", "1")
	q.Add("convert", "UAH")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", key)

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func unmarshalResponse(body []byte) (BTC, error) {
	var bit BTC
	err := json.Unmarshal(body, &bit)
	if err != nil {
		return bit, err
	}

	return bit, nil
}
