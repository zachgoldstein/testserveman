package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"math/rand"
	"errors"
)

func main() {
	http.HandleFunc("/test/success", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write( []byte(responseDocument) )

		duration := time.Millisecond * time.Duration(3000 * rand.Float64())
		fmt.Println("sleeping for ",duration)
		time.Sleep(duration)

		if (err != nil){
			fmt.Printf("An error occurred writing the document %v", err)
		}

		fmt.Printf("Returned a doc %v", responseDocument)
	})

	http.HandleFunc("/test/fail/validate", func(w http.ResponseWriter, r *http.Request) {
		errorChance := 0.50
		var err error
		if errorChance > rand.Float64() {
			_, err = w.Write( []byte(responseDocument) )
		} else {
			_, err = w.Write( []byte(failResponseDocument) )
			return
		}
		if (err != nil){
			fmt.Errorf("An error occurred writing the document %v", err)
		}

		duration := time.Millisecond * time.Duration(3000 * rand.Float64())
		fmt.Println("sleeping for ",duration)
		time.Sleep(duration)

		fmt.Printf("Returned a doc %v", responseDocument)
	})

	http.HandleFunc("/test/fail/error", func(w http.ResponseWriter, r *http.Request) {
		var err error
		err = errors.New("Testing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Errorf("A testing error is returing %v", err)

		duration := time.Millisecond * time.Duration(3000 * rand.Float64())
		fmt.Println("sleeping for ",duration)
		time.Sleep(duration)

		fmt.Printf("Returned a doc %v", responseDocument)
	})

	http.HandleFunc("/test/fail/mix", func(w http.ResponseWriter, r *http.Request) {
		errorChance := 0.50
		var err error
		if errorChance > rand.Float64() {
			err = errors.New("Testing error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Errorf("A testing error is returing %v", err)
			return
		} else {
			_, err = w.Write( []byte(failResponseDocument) )
			fmt.Errorf("An error occurred writing the document %v", err)
			return
		}

		duration := time.Millisecond * time.Duration(3000 * rand.Float64())
		fmt.Println("sleeping for ",duration)
		time.Sleep(duration)

		fmt.Printf("Returned a doc %v", responseDocument)
	})

	fmt.Printf("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

var failResponseDocument = `
{
    "id": 1,
    "name": "A green door",
    "stringNumber": "wtf",
    "price": "this should never happen!!!!",
    "tags": ["home", "green"]
}
`
