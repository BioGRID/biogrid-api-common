// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package requests

import (
	"net/http"
	"encoding/json"
	"errors"
	"github.com/BioGRID/biogrid-api-common/validation"
)

func ProcessBody( r *http.Request, d interface{}, v *validation.ValidationHandler ) ([]string,error) {

	issues := []string{}

	// Check body exists
	if r.Body == nil {
		return issues, errors.New( "Unable to process request without body." )
	}

	// Pull data out of JSON request body
	err := decodeJSONBody( r, d)
	if err != nil {
		return issues, errors.New( "Incorrectly formatted json in request. Check specifications for correct JSON request body." )
	}

	// Perform validation
	if v != nil {
		issues := v.ValidateStruct(d)
		if len(issues) > 0 {
			return issues, errors.New( "Request failed validation." )
		}
	}

	return issues, nil
}

// Decode the json body of a request
func decodeJSONBody( r *http.Request, d interface{} ) (error) {
	decoder := json.NewDecoder( r.Body )
	decoder.DisallowUnknownFields( )
	err := decoder.Decode( &d )
	return err
}