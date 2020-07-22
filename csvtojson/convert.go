package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

//Book ...
type Book struct {
	ID       string
	Name     string
	Author   string
	Rating   string
	ISBN10   string
	ISBN13   string
	Language string
	Pages    string
	Ratings  string
	Reviews  string
}

var books []Book

func main() {

	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatal(err)
	}

	rdr := csv.NewReader(csvFile)

	allData, err := rdr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	book := Book{"", "", "", "", "", "", "", "", "", ""}
	for _, v := range allData {

		for j, b := range v {
			reflect.ValueOf(&book).Elem().Field(j).SetString(b)
		}
		books = append(books, book)
	}

	byteData, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("books.json", byteData, 0644)
}
