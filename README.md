
# Heroku Ip Spoofer
Heroku ip spoofer by tristan
// readme not finished yet
// code not fully optimized yet
<br>

# Usage
Login to Heroku CLI
```
$ heroku login
```

<br>

Commit the api files in /heroku
```
$ git init
$ git add *
$ git commit -m "commit"
```

<br>

Create Heroku Application
```
$ heroku create
```

# Heroku Api Usage (Spoof IP)
```go
package main

// Import the required modules
import (
	"encoding/base64"
	"fmt"

	// Fasthttp for sending the request to api
	"github.com/valyala/fasthttp"
)

// Base64 Encode
func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Function to test the localhost endpoint
func main() {
	// Create new fasthttp request and response object
	var (
        // Fasthttp request client
        requestClient *fasthttp.Client = &fasthttp.Client{}

        // Fasthttp request object
		req        *fasthttp.Request  = fasthttp.AcquireRequest()

        // Fasthttp response object
		resp       *fasthttp.Response = fasthttp.AcquireResponse()

        // Base64 Encoded url
		encodedUrl string             = Base64Encode([]byte("http://httpbin.org/"))

        // Your heroku api url
        herokuApiUrl string = "Your heroku api url"
	)
	// Release the request once no longer needed
	defer fasthttp.ReleaseRequest(req)

	// Set the request url to localhost with the encodedurl param
	req.Header.SetRequestURI(fmt.Sprintf("%s?url=%s", herokuApiUrl, encodedUrl))

	// Set the request method (GET)
	req.Header.SetMethod("GET")

	// Sent the http request
	requestClient.Do(req, resp)

	// Print response body
	fmt.Println(string(resp.Body()))

	// Print response status code
	fmt.Println(resp.StatusCode())
}
```
