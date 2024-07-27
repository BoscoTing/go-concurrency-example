package apiutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchData(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close() // Close the response body when the function returns

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var data []string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return data, nil
}
