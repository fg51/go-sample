package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Fruit is JSON struct
type Fruit struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	fromByteToJSON()
	writeJSON()
}

func fromByteToJSON() {
	// bytes, err := ioutil.ReadFile("src.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	bytes := ([]byte)(`[
		{
			"_id": "01",
			"name": "apple",
			"price": 100
		},
		{
			"_id": "02",
			"name": "banana",
			"price": 200
		}
	]`)

	var fruits []Fruit
	if err := json.Unmarshal(bytes, &fruits); err != nil {
		log.Fatal(err)
	}

	for _, f := range fruits {
		fmt.Printf("%s: %d\n", f.Name, f.Price)
	}
}

// Country is JSON structure
type Country struct {
	Name        string       `json:"name"`
	Prefectures []Prefecture `json:"prefectures"`
}

// Prefecture is JSON structure
type Prefecture struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population int    `json:"population"`
}

func writeJSON() {
	cityA := Prefecture{Name: "a", Capital: "A", Population: 1}
	cityB := Prefecture{Name: "b", Capital: "B", Population: 2}
	cityC := Prefecture{Name: "c", Capital: "C", Population: 3}

	country := Country{
		Name:        "a country",
		Prefectures: []Prefecture{cityA, cityB, cityC},
	}

	jsonBytes, err := json.Marshal(country)
	if err != nil {
		log.Fatal(err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "  ")
	fmt.Println(out.String())
}
