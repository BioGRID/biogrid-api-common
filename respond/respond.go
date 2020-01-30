// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package respond

import (
	"encoding/json"
	"net/http"
)

type JSONErrorResponse struct {
	Message     string    	`json:"message"`
	Status      int       	`json:"status"`
	Detail		string		`json:"detail,omitempty"`
	Issues		[]string  	`json:"issues,omitempty"`
}

type JSONSuccessResponse struct {
	Data		interface{}  `json:"data,omitempty"`
}

type JSONErrorDataResponse struct {
	Message     string    	`json:"message"`
	Status      int       	`json:"status"`
	Detail		string		`json:"detail,omitempty"`
	Data		interface{}	`json:"data,omitempty"`
}

func RESPOK( w http.ResponseWriter ) {
	w.WriteHeader(http.StatusOK)
}

// Shortcut to respond with a 200 status code
func JSONOK( w http.ResponseWriter, data interface{} ) {
	JSONCode( w, http.StatusOK, data )
}

// Shortcut to respond with a status code and wrapped data
func JSONData( w http.ResponseWriter, status int, data interface{} ) {
	resp := JSONSuccessResponse{ Data: data }
	JSONCode( w, status, resp )
}

// Format result as an error message response and then send
// with appropriate error code
func JSONError( w http.ResponseWriter, status int, message string ) {
	err := JSONErrorResponse{ Status: status, Message: message }
	JSONCode( w, status, err )
}

// Format result as an error message response and then send
// with appropriate error code
func JSONErrorWithIssues( w http.ResponseWriter, status int, message string, issues []string ) {
	err := JSONErrorResponse{ Status: status, Message: message, Issues: issues }
	JSONCode( w, status, err )
}

// Format response header and encode interface
// for standardized json response
func JSONCode( w http.ResponseWriter, status int, data interface{} ) {
	w.Header( ).Set( "Content-Type", "application/json" )
	w.Header( ).Set( "Access-Control-Allow-Origin", "*" )
	//w.Header( ).Set( "Connection", "close" )
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Format result as an error message response and then send
// with appropriate error code and data packet
func JSONErrorWithData( w http.ResponseWriter, status int, message string, data interface{}, detail string ) {
	resp := JSONErrorDataResponse{ Status: status, Message: message, Data: data }
	if detail != "" {
		resp = JSONErrorDataResponse{ Status: status, Message: message, Detail: detail, Data: data }
	}
	JSONCode( w, status, resp )
}

// Format Result as an error message response with an additional detail field
func JSONErrorWithDetail( w http.ResponseWriter, status int, message string, detail string ) {
	resp := JSONErrorResponse{ Status: status, Message: message }
	if detail != "" {
		resp = JSONErrorResponse{ Status: status, Message: message, Detail: detail }
	}
	JSONCode( w, status, resp )
}

// Output results as Bytes
func BYTEData( w http.ResponseWriter, status int, data []byte ) {
	w.Header( ).Set( "Content-Type", "application/json" )
	w.Header( ).Set( "Access-Control-Allow-Origin", "*" )
	w.WriteHeader(status)
	w.Write(data)
}