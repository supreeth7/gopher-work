package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

func main() {
	jsonFile := flag.String("file", "quiz.json", "The JSON file with the quiz questions. Expected format: 'question:answer'")
	flag.Parse()

	file, err := os.Open(*jsonFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		os.Exit(1)
	}

	fileName, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	q := []quiz{}

	err = json.Unmarshal(fileName, &q)
	if err != nil {
		fmt.Printf("Failed to unmarshal file: %v", err)
		os.Exit(1)
	}

	var score int

	reader := bufio.NewReader(os.Stdin)

	for i, question := range q {
		fmt.Printf("#%d %s?\n", i+1, question.Question)

		if scrutinize(question.Answer, reader) {
			score++
		}
	}

	fmt.Printf(
		"****************\nTotal score: %d/%d\n****************\n",
		score,
		len(q),
	)
}

func scrutinize(answer string, reader *bufio.Reader) bool {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v", err)
		os.Exit(1)
	}

	input = strings.Trim(input, "\n")

	if answer == input {
		fmt.Println("Correct!")
		return true
	}

	fmt.Println("Better luck next time!")
	return false
}
