// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package requests_test

import (
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/BioGRID/biogrid-api-common/requests"
	"github.com/BioGRID/biogrid-api-common/testutils"
	"github.com/BioGRID/biogrid-api-common/validation"
)

var vh validation.ValidationHandler

func init( ) {
	vh.Initialize( )
}

func TestRequests_NilRequestBody( t *testing.T ) {

	var test = struct{
		param	string
	}{
		"test",
	}

	request, _ := http.NewRequest( "", "", nil )
	issues, err := requests.ProcessBody( request, &test, nil )
	assert.Equal( t, len(issues), 0 )
	assert.NotNil( t, err )
	assert.Equal( t, err.Error( ), "Unable to process request without body." )

}

type ValidReq struct {
	Field string `json:"field" validate:"printascii,required,min=10"`
}

func TestRegister_JSONBodyFormatting( t *testing.T ) {

	var testVal = ValidReq{}

	var tests = []struct{
		param     string
		expectNil bool
		note      string
	} {
		{``, false,"Blank"},
		{`{"field":1}`, false, "Valid JSON/InValid Body"},
		{`{"field":"test"}`, true, "Valid JSON/Valid Body"},
		{`{"field":""}`, true, "Valid JSON/Empty Field"},
		{`{"id":1}`, false, "Valid JSON/InValid Body"},
		{`{}`, true, "Valid JSON/Empty Body"},
		{`{""}`, false, "InValid JSON/Empty Body"},
	}

	for _, test := range tests {
		testutils.OutputTestNote( t, test.note )
		r, _ := http.NewRequest( "POST", "", bytes.NewBufferString(test.param) )
		issues, err := requests.ProcessBody( r, &testVal, nil )
		assert.Equal( t, len(issues), 0 )
		if test.expectNil {
			assert.Nil( t, err )
		} else {
			assert.NotNil( t, err )
			if err != nil {
				assert.Equal( t, err.Error( ), "Incorrectly formatted json in request. Check specifications for correct JSON request body." )
			}
		}
		
	}

}

func TestRegister_JSONBodyValidation( t *testing.T ) {

	var testVal = ValidReq{}

	var tests = []struct{
		param     string
		expectNil bool
		issues    int
		note      string
	} {
		{`{"field":"test"}`, false, 1, "String too short"},
		{`{"field":""}`, false, 1, "String empty"},
		{`{}`, false, 1, "String missing"},
		{`{"field":"testtesttest"}`, true, 0, "String is valid"},
	}

	for _, test := range tests {
		testutils.OutputTestNote( t, test.note )
		r, _ := http.NewRequest( "POST", "", bytes.NewBufferString(test.param) )
		issues, err := requests.ProcessBody( r, &testVal, &vh )
		assert.Equal( t, test.issues, len(issues) )
		if test.expectNil {
			assert.Nil( t, err )
		} else {
			assert.NotNil( t, err )
			if err != nil {
				assert.Equal( t, err.Error( ), "Request failed validation." )
			}
		}
		
	}

}