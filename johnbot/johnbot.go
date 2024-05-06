package johnbot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type BotResponse struct {
	Greetings []string          `json:"greetings"`
	Questions []string          `json:"questions"`
	Yes       map[string]string `json:"yes"`
	No        []string          `json:"no"`
	Default   []string          `json:"default"`
	Farewells []string          `json:"farewells"`
}

const (
	questions = "questions"
)

func (r *BotResponse) LoadBotResponse(filename string) error {
	// Read the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error when opening the file: %v", err)
	}
	defer file.Close()

	// Parse the JSON data
	err = json.NewDecoder(file).Decode(&r)
	for _, greeting := range r.Greetings {
		fmt.Println(greeting)
	}
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	// Access and print the greetings

	return nil
}

func (r *BotResponse) LoadBotQuestions(filename string) error {
	// Open the JSON file
	bye := regexp.MustCompile(`(?i)\b(bye|goodbye|farewell|see you|see ya|take care)\b`)
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	defer file.Close()

	// Decode the JSON file into BotResponse struct
	if err := json.NewDecoder(file).Decode(&r); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Iterate over questions and prompt the user
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; ; i++ {
		questionKey := fmt.Sprintf("%d", i)
		question, ok := r.Yes[questionKey]
		if !ok {
			break // No more questions
		}
		fmt.Println(question)

		// Get user input
		scanner.Scan()
		userInput := strings.ToLower(scanner.Text())
		if bye.MatchString(userInput) {
			r.LoadBoatGoodbyes("johnbot.json")
		}
		if len(userInput) > 10 {
			fmt.Println("OK, next question")
		}
	}

	return nil
}

func (r *BotResponse) LoadBoatGoodbyes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	defer file.Close()

	// Decode the JSON file into BotResponse struct

	err = json.NewDecoder(file).Decode(&r)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}
	for _, farewell := range r.Farewells {
		fmt.Println(farewell)
		os.Exit(0)

	}

	return nil
}

func ReadUserInput() {
	var r BotResponse
	bye := regexp.MustCompile(`(?i)\b(bye|goodbye|farewell|see you|see ya|take care)\b`)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to greenbelt bot")
	for {
		scanner.Scan()
		userInput := strings.ToLower(scanner.Text())

		if userInput == "hello" {
			r.LoadBotResponse("johnbot.json")
		}

		if userInput == questions {
			r.LoadBotQuestions("johnbot.json")
		}

		if bye.MatchString(userInput) {
			r.LoadBoatGoodbyes("johnbot.json")
		}
	}
}

func Main() int {

	ReadUserInput()
	return 0
}
