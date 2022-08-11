
# Heroku IP Spoofer
![banner (2)](https://user-images.githubusercontent.com/75189508/184139417-4559d58c-965b-4198-82a4-74b5b889ecbb.png)


Heroku IP Spoofer uses the heroku servers to send an http request through an api. Doing this changes the outgoing http request's ip address.
<br>

# Usage
Login to Heroku CLI
```
$ heroku login
```

Commit the api files in /heroku
```
$ git init
$ git add *
$ git commit -m "commit"
```

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
	var (
		// Fasthttp request client
		requestClient *fasthttp.Client = &fasthttp.Client{}

		// Fasthttp request object
		req *fasthttp.Request  = fasthttp.AcquireRequest()

		// Fasthttp response object
		resp *fasthttp.Response = fasthttp.AcquireResponse()

		// Base64 Encoded url
		encodedUrl string = Base64Encode([]byte("https://api.ipify.org?format=json"))

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
