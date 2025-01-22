package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchAPIData(functionName, symbol, apiKey string) ([]byte, error) {
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