package makeRequest

import (
	"bytes"
	"fmt"
	"net/http"
)

type RequestOptionsProps struct {
	Method  string
	APIPath string
	Body    []byte
}

type Header struct {
	Prop string
	Val  string
}

func MakeRequest(options RequestOptionsProps, headers []Header) *http.Response {
	base := "https://www.avanza.se"

	fmt.Printf(`Calling %s`, base+options.APIPath)

	req, _ := http.NewRequest(options.Method, base+options.APIPath, bytes.NewBuffer(options.Body))

	// Add relevant headers to the request
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	for _, headerElement := range headers {
		req.Header.Set(headerElement.Prop, headerElement.Val)
	}

	client := http.Client{}

	res, _ := client.Do(req)

	return res

}
