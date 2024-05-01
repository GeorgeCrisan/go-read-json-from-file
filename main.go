package main

import (
	"fmt"
	"json-playground/utils"
)

func main() {
	rawData := utils.ReadJsonFromFile("people.json")
	result := utils.GetPersonsCollection(rawData, true)
	fmt.Printf("Result length: %v\n", len(result))
}
