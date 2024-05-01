package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Friend struct {
	Name string `json:"name"`
}

type Person struct {
	Name          string `json:"name"`
	Gender        string
	Email         string
	BestFriends   []Friend
	FavoriteFruit string
}

func check(err error) {
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}
}

func ReadJsonFromFile(fileName string) []byte {

	dataByes, err := os.ReadFile(fileName)

	check(err)

	return dataByes
}

func GetPersonsCollection(data []byte, withPrintSample bool) []Person {
	var collection []Person

	err := json.Unmarshal(data, &collection)

	check(err)

	if withPrintSample && len(collection) > 0 {
		sample, err := json.MarshalIndent(collection[0], "", "\t")
		check(err)
		fmt.Printf("Here is a sample: %+v\n", string(sample))
	}

	return collection
}
