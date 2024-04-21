package main

import (
	"net/http"
)

// AddCustomHeader adds custom "x-custom-header: Timedotcom Berhad" header to the request
func AddCustomHeader(rw http.ResponseWriter, r *http.Request) {
	logger.Info("Processing HTTP request in Golang plugin!!")
	r.Header.Add("x-custom-header", "Timedotcom Berhad")
}

func main() {}
