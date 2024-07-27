package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"goroutine/common/apiutil"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL")
		os.Exit(1)
	}
	url := os.Args[1]

	times := 10
	if len(os.Args) > 2 {
		count, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid count value")
			os.Exit(1)
		}
		times = count
	}
	start := time.Now()
	succCount := 0
	errCount := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex	
	for i := 0; i < times; i++ {
		wg.Add(1) // Add 1 to the WaitGroup counter

		go func() {
			defer wg.Done()
			_, err := apiutil.FetchData(url)
			mutex.Lock()
			if err != nil {
				errCount++
			} else {
				succCount++
			}
			mutex.Unlock()
		}()
	}
	wg.Wait()
	elasped := time.Since(start)
	fmt.Println("Execution time:", elasped)
	fmt.Println("Error Count:", errCount)
	succRate := float64(succCount) / float64(times) * 100
	fmt.Println("Success Count:", succCount)
	fmt.Printf("Success Rate: %.2f%%\n", succRate)
}

// func fetchData(url string) ([]string, error) {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("error making request: %v", err)
// 	}
// 	defer response.Body.Close() // Close the response body when the function returns

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading response: %v", err)
// 	}

// 	var data []string
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// 	}
// 	return data, nil
// }
