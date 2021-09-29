package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type SwCharacters struct {
	Name 				string `json:"name"`
	Height 			string `json:"height"`
	Mass 				string `json:"mass"`
	BirthYear 	string `json:"birth_year"`
	Gender 			string `json:"gender"`
	Url 				string `json:"url"`
}

func getCharacter(baseUrl string) (character *SwCharacters, err error) {
	response, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&character)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func getCharactersConcurrently(numOfChars int) {
	var charsMap sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < numOfChars; i++ {
		wg.Add(1)
		go func(index int) {
			character, err := getCharacter("https://swapi.dev/api/people/" + strconv.Itoa(index))
			if err != nil {
				panic(err)
			}
			charsMap.Store(index, character)
			fmt.Printf("New Star Wars character: %v\n", character.Name)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func printExecutionTime(t time.Time) {
	fmt.Println("Execution time: ", time.Since(t))
}

func main() {
	startTime := time.Now()
	defer printExecutionTime(startTime)
	getCharactersConcurrently(100)
}