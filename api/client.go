package api

import (
	"fmt"
	"encoding/csv"
	"io/ioutil"
	"net/http"
)

func FetchAPIData(functionName string, symbol string, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=%s&symbol=%s&apikey=%s", functionName, symbol, apiKey)
	resp, err := http.Get(url)
	
	if err != nil {
		return nil, fmt.Errorf("error fetching data from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected API response status: %d", resp.StatusCode)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("error reading response body: %w", err)
	// }

	// //Print the raw JSON response
	// fmt.Println("Raw JSON response:")
	// fmt.Println(string(body))
	return ioutil.ReadAll(resp.Body)
}

func FetchCSVReader(functionName string, horizon string, apiKey string) (*csv.Reader, func(), error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=%s&horizon=%s&apikey=%s", functionName, horizon, apiKey)
	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("error making GET request: %v", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Return the csv.Reader object
	return csv.NewReader(resp.Body), func() { resp.Body.Close() }, nil
}