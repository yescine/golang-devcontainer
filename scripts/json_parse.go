package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	// Completed bool   `json:"completed"`
	hash string // private field
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var todoItem Todo
		decoder := json.NewDecoder(response.Body)

		//! decoder.DisallowUnknownFields()

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal("Decode error:", err)
		} else {
			fmt.Printf("Data from API: %+v\n", todoItem)
		}

		// Encode to json
		todo, _ := json.MarshalIndent(todoItem, "", "\t")
		fmt.Println("\n\n", string(todo))

	} else {
		log.Fatalf("Request failed with status: %d", response.StatusCode)
	}
}
