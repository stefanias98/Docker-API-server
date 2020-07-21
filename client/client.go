package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//BlockOfData is a struct which hold the relevant data of a block
type blockOfData struct {
	ID     string   `json:"Brand ID"`
	Brand  string   `json:"Manufacturer"`
	Models []string `json:"Current Models"`
	Colors []string `json:"Colors"`
}

func main() {
	if os.Args[1] == "post" {
		postData()
	} else if os.Args[1] == "delete" {
		deleteData(os.Args[2])
	}

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

func deleteData(brand string) {
	fmt.Println("http://localhost:10000/deleteCar/" + brand)

	client := &http.Client{}

	request, err := http.NewRequest("DELETE", "http://localhost:10000/deleteCar/"+brand, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
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
