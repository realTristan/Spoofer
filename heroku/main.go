package main

// Import the required modules
import (
	"encoding/base64"
	"os"

	// Fasthttp to host the api
	"github.com/valyala/fasthttp"
)

// Fasthttp Request Client
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// Request Struct
// The Request structure holds four keys
//   - Method: string -> the request method
//   - Url: string -> the request url
//   - Body: []byte -> the request body
//   - Headers: map[string]string -> the request headers
type Request struct {
	Method  string
	Url     string
	Context *fasthttp.RequestCtx
}

// The Base64Decode() function will decode a base64 encrypted string
func Base64Decode(b []byte) string {
	// Decode the string(b []byte)
	var decoded, _ = base64.StdEncoding.DecodeString(string(b))

	// Return the decoded string
	return string(decoded)
}

// The SetRequest() function will create a new fasthttp request object
// then set it's url and method using the Request struct
//
// After the method and url have been set, it will then set the request headers
// from the Request struct (previously from the GetHeaderMap() function)
func SetRequest(req *Request) *fasthttp.Request {
	// Create new request object
	var request *fasthttp.Request = fasthttp.AcquireRequest()

	// Set request url and method
	request.SetRequestURI(req.Url)
	request.Header.SetMethod(req.Method)

	// Set Request Headers
	req.Context.Request.Header.VisitAll(func(k []byte, v []byte) {
		request.Header.Set(string(k), string(v))
	})
	return request
}

// The SetResponse() function will set the response object
// for the outgoing request
//
// Using the Request struct it will enable / disable the response body
// through the resp.SkipBody fasthttp function
func SetResponse(req *Request) *fasthttp.Response {
	// Create a new fasthttp response object
	var resp *fasthttp.Response = fasthttp.AcquireResponse()

	// Whether to skip the response body if 'HEAD' request
	resp.SkipBody = req.Method == "HEAD"

	// Return response object
	return resp
}

// The SendHttpRequest() function will send the outgoing http request
// to the Request struct url
//
// If the Request struct method is POST or PUT then it will set the
// request body using the defined Request struct body
//
// After the method and body have been set, using the RequestClient
// variable it will send the http request and return the response and the error
func SendHttpRequest(req *Request) (*fasthttp.Response, error) {
	// Define both the fasthttp request and response objects
	var (
		request  *fasthttp.Request  = SetRequest(req)
		response *fasthttp.Response = SetResponse(req)
	)
	// Release the request once no longer needed
	defer fasthttp.ReleaseRequest(request)

	// Add request body if method is POST or PUT
	if req.Method == "POST" || req.Method == "PUT" {
		request.SetBody(req.Context.Request.Body())
	}

	// Send the http request and set the err variables for return
	var err error = RequestClient.Do(request, response)

	// Create Response Object
	return response, err
}

// The SetApiResponse() function will set the localhost api
// response headers, body and status code
func SetApiResponse(ctx *fasthttp.RequestCtx, resp *fasthttp.Response, err error) {
	// Set the response status code
	ctx.SetStatusCode(resp.StatusCode())

	// Set the response body
	ctx.Write(resp.Body())

	// Iterate through the request headers and set the response headers
	resp.Header.VisitAll(func(key []byte, value []byte) {
		ctx.Response.Header.Set(string(key), string(value))
	})
}

// The HandleResponse() function will get the incoming request method
// and Base64Encode the url param
//
// It will then create a new Request struct object with the url, method, body,
// whether to skip the body, and the incoming request headers
//
// Once the Request struct object has been created, it will send the http request
// and handle the response using the SetApiResponse() function
func HandleResponse(ctx *fasthttp.RequestCtx) {
	var (
		// Get the request method
		method string = string(ctx.Request.Header.Method())

		// Get the decoded url to send http request to
		url string = Base64Decode(ctx.QueryArgs().Peek("url"))

		// Create a new Request struct object
		req *Request = &Request{
			Url:     url,
			Method:  method,
			Context: ctx,
		}

		// Get the sent request response and error
		resp, err = SendHttpRequest(req)
	)

	// Handle the sent http request
	SetApiResponse(ctx, resp, err)
}

// Main function
func main() {
	// Get the host port from the .env file
	var port string = os.Getenv("PORT")

	// Listen And Server (Host) the api to the corresponding function
	fasthttp.ListenAndServe(":"+port, func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			HandleResponse(ctx)
		}
	})
}
