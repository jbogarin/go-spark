package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

// PrintRequestWithoutBody prints the request without the Body, mainly for GET requests
func PrintRequestWithoutBody(request *http.Request) {
	fmt.Fprintln(os.Stderr, "Request\nURL:", request.URL)
	fmt.Fprintln(os.Stderr, "Method:", request.Method)
	fmt.Fprintln(os.Stderr, "Headers:")
	for key, value := range request.Header {
		fmt.Fprintln(os.Stderr, "\t"+key+":", value[0])
	}
	fmt.Fprintln(os.Stderr, "\nOutput")
}

// PrintRequestWithBody prints the request with the Body
func PrintRequestWithBody(request *http.Request, body interface{}) {
	fmt.Fprintln(os.Stderr, "Request\nURL:", request.URL)
	fmt.Fprintln(os.Stderr, "Method:", request.Method)
	fmt.Fprintln(os.Stderr, "Headers:")
	for key, value := range request.Header {
		fmt.Fprintln(os.Stderr, "\t"+key+":", value[0])
	}
	bodyJSON, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "Body:\n", string(bodyJSON))
	fmt.Fprintln(os.Stderr, "\nOutput")
}

// PrintJSON prints the response to stdout
func PrintJSON(response interface{}) {
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseJSON))

}

// PrintCSV prints the response in CSV to stdout
func PrintCSV(response interface{}) {
	csvContent, err := gocsv.MarshalString(response) // Get all clients as CSV string
	if err != nil {
		panic(err)
	}
	fmt.Print(csvContent) // Display all clients as CSV string

}

// PrintResponseFormat prints the response depending on the format flag
func PrintResponseFormat(response interface{}) {
	if format == "json" {
		PrintJSON(response)
	} else if format == "csv" {
		PrintCSV(response)
	} else {
		PrintJSON(response)
	}
}
