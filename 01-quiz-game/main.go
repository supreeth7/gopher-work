package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type quiz struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

func main() {
	jsonFile := flag.String("file", "quiz.json", "The JSON file with the quiz questions. Expected format: 'question:answer'")

	limit := flag.Int("limit", 30, "The time limit for the quiz; seconds")

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
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, question := range q {
		fmt.Printf("\n#%d %s?\n", i+1, question.Question)

		answer := make(chan bool)
		go scrutinize(question.Answer, reader, answer)

		select {
		case <-timer.C:
			fmt.Printf(
				"****************\nTotal score: %d/%d\n****************\n",
				score,
				len(q),
			)
			return
		case ans := <-answer:
			if ans {
				fmt.Println("Correct!")
				score++
			}
			fmt.Println("Better luck next time!")
		}
	}
}

func scrutinize(answer string, reader *bufio.Reader, ch chan bool) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v", err)
		os.Exit(1)
	}

	input = strings.Trim(input, "\n")

	if answer == input {
		ch <- true
	}

	ch <- false
}
