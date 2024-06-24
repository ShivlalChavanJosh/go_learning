package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	word, err := getRandomWord()
	if err != nil {
		fmt.Println("Error fetching the random word:", err)
		return
	}

	maxChances := 8

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
		fmt.Printf("Guess a letter: ")

		// Take the input
		var str string
		fmt.Scanln(&str)
		str = strings.ToLower(strings.TrimSpace(str))

		// Ensure input is a single character
		if len(str) != 1 {
			fmt.Println("Please enter only a single character.")
			continue
		}

		// Skip if the input is already guessed
		if contains(entries, str) {
			fmt.Println("You've already guessed that. Try something else.")
			continue
		}

		// Add the input to entries
		entries = append(entries, str)

		// Compare and update entries, placeholder and chances
		if strings.Contains(word, str) {
			for i := range word {
				if string(word[i]) == str {
					placeholder[i] = str
				}
			}
		} else {
			chancesLeft--
		}
	}
}

func getRandomWord() (string, error) {
	resp, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", fmt.Errorf("no words found")
	}

	fmt.Printf("WORD: ", words[0])

	return words[0], nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
