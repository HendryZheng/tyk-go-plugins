package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// MyPluginResponse intercepts response from upstream
func MyPluginResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) {
	// Add a header to our response object
	res.Header.Add("X-Response-Added", "resp-added")

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(rw, "Error reading response body", http.StatusInternalServerError)
		return
	}

	// Note: Body is a stream and can be read only once and you can already use body variable to read the response body as a string or byte array.
	// Log the original response body
	log.Printf("Original response body: %s", string(body))

	// Close the original response body
	res.Body.Close()

	// TO DO: Add your custom logic here to modify the response body

	// LOGIC TO CONVERT PIPE DELIMITED DATA TO JSON
	// Split the input string by the delimiter "|"
	splitData := bytes.Split([]byte(body), []byte("|"))

	// Creating a map from the split data
	dataMap := make(map[string]string)
	for i, data := range splitData {
		key := "field_" + strconv.Itoa(i+1)
		value := string(data)
		dataMap[key] = value
	}

	// Log the extracted data
	log.Printf("Extracted data: %+v", dataMap)

	// Convert the map to JSON
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		log.Printf("Error marshalling data to JSON: %v", err)
		http.Error(rw, "Error converting data to JSON", http.StatusInternalServerError)
		return
	}

	// Overwrite our response body with the JSON data
	res.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	res.ContentLength = int64(len(jsonData))
	res.Header.Set("Content-Length", strconv.Itoa(len(jsonData)))
	res.Header.Set("Content-Type", "application/json")

	// Log the modified response body
	log.Printf("Modified response body: %s", string(jsonData))
}
