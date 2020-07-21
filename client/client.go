package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//BlockOfData is a struct which hold the relevant data of a block
type blockOfData struct {
	ID     string   `json:"Brand ID"`
	Brand  string   `json:"Manufacturer"`
	Models []string `json:"Current Models"`
	Colors []string `json:"Colors"`
}

func main() {
	postData()
}

func postData() {
	newCarData := blockOfData{ID: "3", Brand: "Porsche", Models: []string{"911 Carrera GTS", "911 Turbo S", "911 GT3 RS", "911 GT2 RS"}, Colors: []string{"GT Silver", "Yachting Blue Metallic", "British Racing Green", "Voodoo Blue", "Oak Green Metallic"}}

	marData, err := json.MarshalIndent(newCarData, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:10000/addCar", "application/json", bytes.NewBuffer(marData))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))

}
