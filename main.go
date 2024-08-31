package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	errWrongInput = "An error occured while reading input. Please try again."
)

func main() {
	var err error
	var reader *bufio.Reader

	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	minLength := 0
	maxLength := 0
	givenLetters := ""

	fmt.Print("Minimal word length: ")
	if _, err = fmt.Scanf("%d\n", &minLength); err != nil {
		fmt.Println(errWrongInput, err)
		return
	}

	fmt.Print("Maximal word length: ")
	if _, err = fmt.Scanf("%d\n", &maxLength); err != nil {
		fmt.Println(errWrongInput, err)
		return
	}

	fmt.Print("Given letters: ")
	reader = bufio.NewReader(os.Stdin)

	givenLetters, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(errWrongInput, err)
		return
	}

	givenLetters = strings.TrimSpace(givenLetters)
	givenLetters = strings.ToLower(givenLetters)

	var possibleWords []string
	for _, word := range words {
		word = strings.ToLower(word)
		wordLength := len(word)
		if wordLength >= minLength && wordLength <= maxLength && constructWord(word, givenLetters) {
			possibleWords = append(possibleWords, word)
		}
	}

	fmt.Printf("Possible words (%d):\n", len(possibleWords))
	for _, word := range possibleWords {
		fmt.Println(word)
	}
}

func constructWord(word string, givenLetters string) bool {
	letterCount := make(map[rune]int)

	for _, letter := range givenLetters {
		letterCount[letter]++
	}

	for _, letter := range word {
		if letterCount[letter] == 0 {
			return false
		}
		letterCount[letter]--
	}
	return true
}
