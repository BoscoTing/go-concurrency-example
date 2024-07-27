package main

import (
	"fmt"
	"os"
	"strconv"
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
	for i := 0; i < times; i++ {
		_, err := apiutil.FetchData(url)
		if err != nil {
			errCount++
		} else {
			succCount++
		}
	}
	elasped := time.Since(start)
	fmt.Println("Execution time:", elasped)
	fmt.Println("Error Count:", errCount)
	succRate := float64(succCount) / float64(times) * 100
	fmt.Println("Success Count:", succCount)
	fmt.Printf("Success Rate: %.2f%%\n", succRate)
}
