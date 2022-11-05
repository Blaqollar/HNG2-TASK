package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type DataType struct {
	OperationType string `json:"operation_type"`
	X             int64  `json:"x"`
	Y             int64  `json:"y"`
}

type Output struct {
	SlackUsername string `json:"slackUsername"`
	OperationType string `json:"operation_type"`
	Result        int64  `json:"result"`
}

func GetProducts(rw http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var d DataType
	var o Output
	json.Unmarshal(reqBody, &d)
	if d.OperationType == "addition" {
		o.Result = d.X + d.Y
	} else if d.OperationType == "subtraction" {
		o.Result = d.X - d.Y
	} else if d.OperationType == "multiplication" {
		o.Result = d.X * d.Y
	}
	o.OperationType = d.OperationType
	o.SlackUsername = "uchedingba"

	json.NewEncoder(rw).Encode(o)
	_, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sm := mux.NewRouter()
	getRouter := sm.Methods("POST").Subrouter()
	getRouter.HandleFunc("/post", GetProducts)

	//Setting up the server //tune the elements to timeout to avoid bad connections errors
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
