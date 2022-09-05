package main

// Import the required modules
import (
	"encoding/base64"
	"fmt"

	// Fasthttp for sending the request to api
	"github.com/valyala/fasthttp"
)

// Fasthttp request client
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// Base64 Encode
func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Function to test the localhost endpoint
func main() {
	// Create new fasthttp request and response object
	var (
		req        *fasthttp.Request  = fasthttp.AcquireRequest()
		resp       *fasthttp.Response = fasthttp.AcquireResponse()
		encodedUrl string             = Base64Encode([]byte("https://api.ipify.org?format=json"))
	)
	// Release the request once no longer needed
	defer fasthttp.ReleaseRequest(req)

	// Set the request url to localhost with the encodedurl param
	req.Header.SetRequestURI(fmt.Sprintf("http://localhost:8080?url=%s", encodedUrl))

	// Set the request method (GET)
	req.Header.SetMethod("GET")

	// Sent the http request
	RequestClient.Do(req, resp)

	// Print response body
	fmt.Println(string(resp.Body()))

	// Print response status code
	fmt.Println(resp.StatusCode())
}

// Notes:
// I wasn't properly able to test this because the wifi where I'm currently at
// is horrendous, if there are any bugs please let me know!
