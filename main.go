package main

import (
	"encoding/json"
	"fmt"
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
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/listCars", returnAllData)
	mux.HandleFunc("/listCars/{name}", returnSingleBlock)
	mux.HandleFunc("/listCars2", returnAllDataMarshal)
	mux.HandleFunc("/addCar", writeNewData).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", mux))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().String()
	fmt.Fprintf(w, "Welcome to the REST API homepage! You accessed it at %+v from the IP address %+v \n", timeNow, readUserIP(r))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "If you wish to view all car data, go to path /data. For specific brand data, search for path /data/<brand name>\n ")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Request header data: %+v", r.Header)
	fmt.Println("Path visited: homePage")
	fmt.Println("Server address:", r.Host)
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: /data -> returnAllData")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(data)
}

func returnAllDataMarshal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: /data2 -> returnAllDataMarshal")
	resp, _ := json.MarshalIndent(data, "", "    ")
	w.Write(resp)

}

func returnSingleBlock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: /data/{name} -> returnSingleBlock")
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
		fmt.Fprintf(w, "The specific brand query for %v does not exist in the database \nThe available models are: ", inputVar)
		for _, dataBlock := range data {
			fmt.Fprintf(w, "%v, ", dataBlock.Brand)
		}

	}

}

func writeNewData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: /data POST -> writeNewData")
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
