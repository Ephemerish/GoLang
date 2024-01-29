package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func containsAlphaAndNum(word string) bool {
	hasAlpha, hasNum := false, false

	for _, char := range word {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
			hasAlpha = true
		} else if '0' <= char && char <= '9' {
			hasNum = true
		}
	}

	return hasAlpha && hasNum
}

func main() {
	startTime := time.Now()

	file, err := os.Open("testcur.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	finalwordCount := 0
	var wg sync.WaitGroup
	ch := make(chan int)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, record := range records {
		wg.Add(1)
		go func(record []string, wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			for _, column := range record {
				// Assuming words are space-separated in the column, adjust if needed
				wordCount := 0
				words := strings.Fields(column)

				for _, word := range words {
					if containsAlphaAndNum(word) {
						wordCount++
					}
				}
				ch <- wordCount
			}
		}(record, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for count := range ch {
		finalwordCount += count
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Number of words with alphabet letters and numbers: %d\n", finalwordCount)
	fmt.Printf("Time elapsed: %s\n", elapsedTime)
}
