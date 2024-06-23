package main

import (
	"fmt"
	"strings"
)

func main() {

	word := "elephant"
	maxChances := 6

	// Initialize the entries and placeholder
	entries := []string{}
	placeholder := make([]string, len(word))
	for i := range placeholder {
		placeholder[i] = "_"
	}

	chancesLeft := maxChances

	for {
		// Check for a win
		if strings.Join(placeholder, "") == word {
			fmt.Println("Congratulations! You've guessed the word:", word)
			break
		}

		// Check for a loss
		if chancesLeft <= 0 {
			fmt.Println("You've run out of chances. The word was:", word)
			break
		}

		// Console display
		fmt.Println("\nCurrent word:", strings.Join(placeholder, " "))
		fmt.Printf("Chances left: %d\n", chancesLeft)
		fmt.Printf("Guessed so far: %v\n", entries)
		fmt.Printf("Guess a letter or the word: ")

		// Take the input
		var str string
		fmt.Scanln(&str)
		str = strings.ToLower(str)

		// Skip if the input is already guessed
		if contains(entries, str) {
			fmt.Println("You've already guessed that. Try something else.")
			continue
		}

		// Add the input to entries
		entries = append(entries, str)

		// Compare and update entries, placeholder and chances
		if len(str) == 1 {
			if strings.Contains(word, str) {
				for i := range word {
					if string(word[i]) == str {
						placeholder[i] = str
					}
				}
			} else {
				chancesLeft--
			}
		} else if str == word {
			fmt.Println("Congratulations! You've guessed the word:", word)
			break
		} else {
			chancesLeft--
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
