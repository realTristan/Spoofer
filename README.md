
# Heroku IP Spoofer ![Stars](https://img.shields.io/github/stars/realTristan/Spoofer?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Spoofer?label=Watchers)
![banner (1)](https://user-images.githubusercontent.com/75189508/192170623-1d0fe3c7-bd11-4b30-a01d-a5a24a145ec3.png)

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

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
