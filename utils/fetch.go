package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wvtrserv/logger"
)

func ReadResponse(response *http.Response) []byte {
	// Read and print response
	logger.DumpLog.Println(response)
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}
	}
	//logger.DumpLog.Println(string(resp))
	return resp
}

func Fetch(reqURL string, method string, params url.Values, header []string) *http.Response {
	// Create a new HTTP client
	client := &http.Client{
		Timeout: time.Second * 10, // Timeout each requests
	}

	req := CreateRequest(reqURL, method, params, header)

	//Execute the request using the custom HTTP client
	response, err := client.Do(req)
	if err != nil {
		logger.ErrLog.Println("Error making request:", err)
		return nil
	}

	return response
}

func CreateRequest(reqURL string, method string, params url.Values, header []string) *http.Request {

	logger.DumpLog.Println(reqURL)
	logger.DumpLog.Println(header)
	logger.DumpLog.Println(method)
	logger.DumpLog.Println(params.Encode())

	req, err := http.NewRequest(method, reqURL, strings.NewReader(params.Encode()))

	//logger.DumpLog.Println("request body : ", req.Body)

	if err != nil {
		logger.ErrLog.Println("Error creating request:", err)
		return nil
	}

	if len(header)%2 != 0 {
		logger.ErrLog.Println("Error creating header: the number of parameters don't match")
		return nil
	}

	// Set headers
	for i := 0; i < len(header); i += 2 {
		req.Header.Add(header[i], header[i+1])
	}

	return req
}
