package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//BlockOfData is a struct which hold the relevant data of a block
type BlockOfData struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var data []BlockOfData

func main() {

	data = []BlockOfData{
		BlockOfData{Title: "12 rules for life", Desc: "Whole lotta lobsters", Content: "10^10 lobsters"},
		BlockOfData{Title: "Ben Shapiro and the 3rd trimester", Desc: "My wife is a doctor", Content: "Facts don't care about your pregnancy"},
	}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", returnAllData)
	http.HandleFunc("/data2", returnAllDataMarshal)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().String()
	fmt.Fprintf(w, "Welcome to the HomePage! You accessed it at %+v from the IP address %+v \n", timeNow, ReadUserIP(r))
	fmt.Fprintf(w, "Request header data: %+v", r.Header)
	fmt.Println("Path visited: homePage")
	fmt.Println("Server address:", r.Host)
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: returnAllData")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(data)

}

func returnAllDataMarshal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path visited: returnAllDataMarshal")
	resp, _ := json.MarshalIndent(data, "", "    ")
	w.Write(resp)

}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
