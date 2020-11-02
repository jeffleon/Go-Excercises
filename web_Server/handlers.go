package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API EndPoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(response)

}
