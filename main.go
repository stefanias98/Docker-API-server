package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//BlockOfData is a struct which hold the relevant data of a block
type BlockOfData struct {
	ID     string   `json:"Brand ID"`
	Brand  string   `json:"Manufacturer"`
	Models []string `json:"Current Models"`
	Colors []string `json:"Colors"`
}

var data []BlockOfData

func main() {

	data = []BlockOfData{
		BlockOfData{ID: "1", Brand: "Ferrari", Models: []string{"SF90 Stradale", "812 Superfast", "812 GTS", "GTC4 Lusso", "F8 Tributo", "F8 Spider", "Roma", "Portofino"}, Colors: []string{"Rosso Corsa", "Nero Daytona", "Giallo Modena", "Grigio Silverstone", "Verde Inglese"}},
		BlockOfData{ID: "2", Brand: "Lamborghini", Models: []string{"Aventador SVJ", "Huracan EVO", "Urus"}, Colors: []string{"Verde Hydra", "Rosso Mars", "Verde Singh", "Viola Parsifae"}},
	}
	handleRequests()
}

func handleRequests() {
	mux := mux.NewRouter().StrictSlash(true)

	mux1 := mux.PathPrefix("/api/v1").Subrouter()
	mux1.HandleFunc("/", homePageV1)
	mux1.HandleFunc("/listCars", returnAllData)
	mux1.HandleFunc("/listCars/{name}", returnSingleBlock)
	mux1.HandleFunc("/listCars2", returnAllDataMarshal)
	mux1.HandleFunc("/addCar", writeNewData).Methods("POST")
	mux1.HandleFunc("/deleteCar/{name}", deleteData).Methods("DELETE")

	mux2 := mux.PathPrefix("/api/v2").Subrouter()
	mux2.HandleFunc("/", homePageV2)
	mux2.HandleFunc("/listBooks", listBooks)
	mux2.HandleFunc("/getID/{id}", returnByID)
	mux2.HandleFunc("/getName/{name}", returnByName)
	mux2.HandleFunc("/getAuthor/{author}", returnByAuthor)
	mux2.HandleFunc("/getISBN10/{isbn10}", returnByISBN10)
	mux2.HandleFunc("/getLanguage/{language}", returnByLanguage)

	log.Fatal(http.ListenAndServe(":10000", mux))
}

/*
//////////////////////////////////////////////////////////////////
This is the localhost:10000/api/v1 (mux1) subroute path definition
//////////////////////////////////////////////////////////////////
*/
func homePageV1(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().String()
	fmt.Fprintf(w, "Welcome to the REST API homepage! You accessed it at %+v from the IP address %+v \n", timeNow, readUserIP(r))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "If you wish to view all car data, go to path /listCars. For specific brand data, search for path /listCars/<brand name>\n ")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Request header data: %+v", r.Header)
	fmt.Println("Path visited: homePage")
	fmt.Println("Server address:", r.Host)
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnAllData")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(data)
}

func returnAllDataMarshal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnAllDataMarshal")
	resp, _ := json.MarshalIndent(data, "", "    ")
	w.Write(resp)

}

func returnSingleBlock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnSingleBlock")
	inputVar := mux.Vars(r)["name"]
	found := false

	for _, dataBlock := range data {
		if dataBlock.Brand == inputVar {
			fmt.Fprintf(w, "The specific brand query for %v is below \n", inputVar)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific brand query for %v does not exist in the database \nThe available brands are: ", inputVar)
		for _, dataBlock := range data {
			fmt.Fprintf(w, "%v, ", dataBlock.Brand)
		}

	}

}

func writeNewData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " POST -> writeNewData")

	requestBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "The data posted contains %+v \n", string(requestBody))

	var dataBlock BlockOfData
	err := json.Unmarshal(requestBody, &dataBlock)
	if err != nil {
		log.Fatalf("Most likely encountered incompatible struct format. Error is %+v \n", err)
	}

	data = append(data, dataBlock)
	fmt.Fprintf(w, "Successfuly wrote data")

}

func deleteData(w http.ResponseWriter, r *http.Request) {
	inputVar := mux.Vars(r)["name"]
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " DELETE -> deleteData")

	found := false
	for i, dataBlock := range data {
		if dataBlock.Brand == inputVar {
			data = append(data[:i], data[i+1:]...)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "Unable to delete the brand %v since it is not available in the database \nThe available brands are: ", inputVar)
		for _, dataBlock := range data {
			fmt.Fprintf(w, "%v, ", dataBlock.Brand)
		}

	}

}

/*
//////////////////////////////////////////////////////////////////
This is the localhost:10000/api/v2 (mux2) subroute path definition
//////////////////////////////////////////////////////////////////
*/

func homePageV2(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().String()
	fmt.Fprintf(w, "Welcome to the REST API homepage! You accessed it at %+v from the IP address %+v \n", timeNow, readUserIP(r))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "This is where book data is stored from https://github.com/moficodes/bookdata-api \n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Request header data: %+v", r.Header)
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " ")
	fmt.Println("Server address:", r.Host)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> listBooks")
	jsonFile, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonFile)

}

func returnByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnByID")
	inputVar := mux.Vars(r)["id"]
	found := false
	var jsonMaps []map[string]string

	jsonData, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &jsonMaps)
	for _, dataBlock := range jsonMaps {
		if dataBlock["ID"] == inputVar {
			fmt.Fprintf(w, "The specific ID query for %v is below \n", inputVar)
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific ID query for %v does not exist in the database \n", inputVar)

	}

}

func returnByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnByName")
	inputVar := mux.Vars(r)["name"]
	found := false
	var jsonMaps []map[string]string

	jsonData, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &jsonMaps)
	for _, dataBlock := range jsonMaps {
		if dataBlock["Name"] == inputVar {
			if found == false {
				fmt.Fprintf(w, "The specific Name query for %v is below \n", inputVar)
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific Name query for %v does not exist in the database \n", inputVar)

	}

}

func returnByAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnByAuthor")
	inputVar := mux.Vars(r)["author"]
	found := false
	var jsonMaps []map[string]string

	jsonData, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &jsonMaps)
	for _, dataBlock := range jsonMaps {
		if dataBlock["Author"] == inputVar {
			if found == false {
				fmt.Fprintf(w, "The specific Author query for %v is below \n", inputVar)
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific Author query for %v does not exist in the database \n", inputVar)

	}

}

func returnByISBN10(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnByISBN10")
	inputVar := mux.Vars(r)["isbn10"]
	found := false
	var jsonMaps []map[string]string

	jsonData, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &jsonMaps)
	for _, dataBlock := range jsonMaps {
		if dataBlock["ISBN10"] == inputVar {
			if found == false {
				fmt.Fprintf(w, "The specific ISBN10 query for %v is below \n", inputVar)
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific ISBN10 query for %v does not exist in the database \n", inputVar)

	}

}

func returnByLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: " + html.EscapeString(r.URL.Path) + " -> returnByLanguage")
	inputVar := mux.Vars(r)["language"]
	found := false
	var jsonMaps []map[string]string

	jsonData, err := ioutil.ReadFile("csvtojson/books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &jsonMaps)
	for _, dataBlock := range jsonMaps {
		if dataBlock["Language"] == inputVar {
			if found == false {
				fmt.Fprintf(w, "The specific Language query for %v is below \n", inputVar)
			}
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "    ")
			encoder.Encode(dataBlock)
			found = true
		}
	}

	if found == false {
		fmt.Fprintf(w, "The specific Language query for %v does not exist in the database \n", inputVar)

	}

}

/*
//////////////////////////////////////////////////////////////////
Helper functions
//////////////////////////////////////////////////////////////////
*/
func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
