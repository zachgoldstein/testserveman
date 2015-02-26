package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"math/rand"
	"errors"
	"runtime"
)

var sleepTime = 1000

func main() {
	runtime.GOMAXPROCS(4)

	http.HandleFunc("/test/success", func(w http.ResponseWriter, r *http.Request) {
		writeSuccess(w, r)
	})

	http.HandleFunc("/test/fail/validate", func(w http.ResponseWriter, r *http.Request) {
		writeValidationErr(w, r)
	})

	http.HandleFunc("/test/fail/error", func(w http.ResponseWriter, r *http.Request) {
		writeInternalErr(w, r)
	})

	http.HandleFunc("/test/fail/mix", func(w http.ResponseWriter, r *http.Request) {
		errorChance := 0.33
		if errorChance > rand.Float64() {
			writeSuccess(w, r)
		} else if errorChance * 2 > rand.Float64() {
			writeInternalErr(w, r)
		} else {
			writeValidationErr(w, r)
		}
	})

	fmt.Printf("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeSuccess(w http.ResponseWriter, r *http.Request) {
	duration := time.Millisecond * time.Duration(float64(sleepTime) * rand.Float64())
	fmt.Println("sleeping for ",duration, " before writing doc")
	time.Sleep(duration)

	splitChance := 0.50
	respDoc := ""
	if splitChance > rand.Float64() {
		respDoc = responseDocument
	} else {
		respDoc = largeResponseDocument
	}
	_, err := w.Write( []byte(respDoc) )
	if (err != nil){
		fmt.Printf("An error occurred writing the document %v", err)
	}

	fmt.Printf("%v Returned a doc %v", time.Now(), respDoc)
}

func writeValidationErr(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write( []byte(failResponseDocument) )
	if (err != nil){
		fmt.Errorf("An error occurred writing the document %v", err)
	}

	duration := time.Millisecond * time.Duration(float64(sleepTime) * rand.Float64())
	fmt.Println("sleeping for ",duration)
	time.Sleep(duration)

	fmt.Printf("%v Returned a schema failing doc %v", time.Now(), failResponseDocument)
}

func writeInternalErr(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Error 500 - Testing")
	http.Error(w, err.Error(), http.StatusInternalServerError)

	duration := time.Millisecond * time.Duration(float64(sleepTime) * rand.Float64())
	fmt.Println("sleeping for ",duration)
	time.Sleep(duration)

	fmt.Printf("%v Returned a 500 err", time.Now())
}


var responseDocument = `
{
    "id": 1,
    "name": "A green door",
    "stringNumber": "3",
    "price": 12.50,
    "tags": ["home", "green"]
}
`

var largeResponseDocument = `
{
    "id": 1,
    "name": "Your bones don't break, mine do. That's clear. Your cells react to bacteria and viruses differently than mine. You don't get sick, I do. That's also clear. But for some reason, you and I react the exact same way to water. We swallow it too fast, we choke. We get some in our lungs, we drown. However unreal it may seem, we are connected, you and I. We're on the same curve, just on opposite ends.",
    "stringNumber": "3",
    "price": 12.50,
    "tags": ["home", "green"]
}
`

var failResponseDocument = `
{
    "id": 1,
    "name": "A green door",
    "stringNumber": "wtf",
    "price": "this should never happen!!!!",
    "tags": ["home", "green"]
}
`
