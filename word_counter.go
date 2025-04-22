package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// This function reads a text file and counts how often each word appears:
func getWordCounts(filePath string) (map[string]int, error) {
	// Try to open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()



	// A map to keep track of word appearances:
	counts := make(map[string]int)



	// Use a scanner to go through the file line by line:
	scanner := bufio.NewScanner(file)



	for scanner.Scan() {
		// Break each line into words
		words := strings.Fields(scanner.Text())

		// Go through the words and count them:
		for _, word := range words {
			word = strings.ToLower(word)
			counts[word]++
		}
	}



	// If there was a scanning error, return it:
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}



// This function displays the top 5 most common words:
func showTopWords(counts map[string]int) {
	type wordPair struct {
		Word  string
		Total int
	}

	var pairs []wordPair
	for word, total := range counts {
		pairs = append(pairs, wordPair{Word: word, Total: total})
	}



	// Sort by the total count in descending order:
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Total > pairs[j].Total
	})

	// Show the top 5:
	fmt.Println("Top 5 most used words:")
	for i := 0; i < 5 && i < len(pairs); i++ {
		fmt.Printf("%d. %s: %d\n", i+1, pairs[i].Word, pairs[i].Total)
	}
}

func main() {
	// Ask the user for the file location:
	fmt.Print("Enter the path to the .txt file: ")
	var path string
	fmt.Scanln(&path)

	// Get word frequency:
	counts, err := getWordCounts(path)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}

	showTopWords(counts)
}
