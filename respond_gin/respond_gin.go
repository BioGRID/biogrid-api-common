// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package respond_gin

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
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

func RESPOK( c *gin.Context ) {
	c.Status(http.StatusOK)
}

// Shortcut to respond with a 200 status code
func JSONOK( c *gin.Context, data interface{} ) {
	c.JSON(http.StatusOK, data)
}

// Shortcut to respond with a status code and wrapped data
func JSONData( c *gin.Context, status int, data interface{} ) {
	resp := JSONSuccessResponse{ Data: data }
	JSONCode( c, status, resp )
}

// Format result as an error message response and then send
// with appropriate error code
func JSONError( c *gin.Context, status int, message string ) {
	err := JSONErrorResponse{ Status: status, Message: message }
	JSONCode( c, status, err )
}

// Format result as an error message response and then send
// with appropriate error code
func JSONErrorWithIssues( c *gin.Context, status int, message string, issues []string ) {
	err := JSONErrorResponse{ Status: status, Message: message, Issues: issues }
	JSONCode( c, status, err )
}

// Format response header and encode interface
// for standardized json response
func JSONCode( c *gin.Context, status int, data interface{} ) {
	c.Header( "Content-Type", "application/json; charset=utf-8" )
	c.Header( "Access-Control-Allow-Origin", "*" )
	c.JSON( status, data )
}

// Format result as an error message response and then send
// with appropriate error code and data packet
func JSONErrorWithData( c *gin.Context, status int, message string, data interface{}, detail string ) {
	resp := JSONErrorDataResponse{ Status: status, Message: message, Data: data }
	if detail != "" {
		resp = JSONErrorDataResponse{ Status: status, Message: message, Detail: detail, Data: data }
	}
	JSONCode( c, status, resp )
}

// Format Result as an error message response with an additional detail field
func JSONErrorWithDetail( c *gin.Context, status int, message string, detail string ) {
	resp := JSONErrorResponse{ Status: status, Message: message }
	if detail != "" {
		resp = JSONErrorResponse{ Status: status, Message: message, Detail: detail }
	}
	JSONCode( c, status, resp )
}

// Shortcut to respond with a status code and wrapped data
func BYTEOK( c *gin.Context, data []byte ) {
	BYTECode( c, http.StatusOK, data )
}

// Shortcut to respond with a status code and wrapped data
func BYTEData( c *gin.Context, status int, data []byte ) {
	BYTECode( c, status, data )
}

// Output results as Bytes
func BYTECode( c *gin.Context, status int, data []byte ) {
	c.Header( "Content-Type", "application/json; charset=utf-8" )
	c.Header( "Access-Control-Allow-Origin", "*" )
	c.JSON( status, data )
}