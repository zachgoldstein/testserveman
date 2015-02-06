package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write( []byte(responseDocument) )

		if (err != nil){
			fmt.Printf("An error occurred writing the document %v", err)
		}

		fmt.Printf("Returned a doc %v", responseDocument)
	})

	fmt.Printf("Serving on port 8080 route /test")
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
