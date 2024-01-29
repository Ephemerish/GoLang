package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	// Use ReadAll to read all records at once
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	wordCount := 0

	for _, record := range records {
		for _, column := range record {
			// Assuming words are space-separated in the column, adjust if needed
			words := strings.Fields(column)

			for _, word := range words {
				if containsAlphaAndNum(word) {
					wordCount++
				}
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Number of words with alphabet letters and numbers: %d\n", wordCount)
	fmt.Printf("Time elapsed: %s\n", elapsedTime)
}
