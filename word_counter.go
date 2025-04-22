package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Function to read the file and count word frequency
func countWords(filePath string) (map[string]int, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a map to store word counts
	wordCount := make(map[string]int)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line in the file
	for scanner.Scan() {
		// Split the line into words
		words := strings.Fields(scanner.Text())

		// Count the words
		for _, word := range words {
			// Convert to lowercase to make the count case-insensitive
			word = strings.ToLower(word)
			wordCount[word]++
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCount, nil
}

// Function to print the top 5 most frequent words
func printTopWords(wordCount map[string]int) {
	// Create a slice to hold word frequency pairs
	type wordFrequency struct {
		Word  string
		Count int
	}

	// Convert the map to a slice of wordFrequency
	var frequencies []wordFrequency
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFrequency{Word: word, Count: count})
	}

	// Sort the slice by word count in descending order
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	// Print the top 5 most frequent words
	fmt.Println("Top 5 most frequent words:")
	for i := 0; i < 5 && i < len(frequencies); i++ {
		fmt.Printf("%d. %s: %d\n", i+1, frequencies[i].Word, frequencies[i].Count)
	}
}

func main() {
	// Prompt user for the file path
	fmt.Print("Enter the path to the .txt file: ")
	var filePath string
	fmt.Scanln(&filePath)

	// Count words in the file
	wordCount, err := countWords(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Print the top 5 words
	printTopWords(wordCount)
}
