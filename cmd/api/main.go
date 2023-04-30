package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/bmi", test)

	log.Println("Starting BMI api server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
