package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	//change the output field to be "message"
	Message string `json:"message"`
	//do not output this field
	Author string `json:"-"`
	//do not output the field if the value is empty
	Date string `json:",omitempty"`
	//convert output to a string and rename "id"
	ID int `json:"id, string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	log.Printf("Server starting on port %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{
		Message: "Hello " + request.Name,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func serveCat(w http.ResponseWriter, r *http.Request) {

}
