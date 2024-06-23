package main

import (
	"fmt"
)

func main() {

	var word1, word2 string

	fmt.Println("Enter the first word:")
	fmt.Scanln(&word1)

	fmt.Println("Enter the second word:")
	fmt.Scanln(&word2)

	if word1 == word2 {
		fmt.Println("The words are exactly same")
	} else {
		fmt.Println("The words are not same!")
	}

	if len(word1) != len(word2) {
		fmt.Println("The length is not same for given words")
	}
}
