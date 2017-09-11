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
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "Hello World!",
		Author:  "Stephen King",
		Id:      10,
	}
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic("Ooops!")
	}
	fmt.Fprint(w, string(data))
}
