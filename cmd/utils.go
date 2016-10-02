package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
