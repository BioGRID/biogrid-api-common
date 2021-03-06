// Copyright 2020 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package paramvalidation_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/BioGRID/biogrid-api-common/paramvalidation"
	"github.com/BioGRID/biogrid-api-common/testutils"
)

func TestParamValidation_BoolParam( t *testing.T ) {

	var tests = []struct{
		testDesc	string
		pVal  		string
		pName		string
		expected  	bool
		isErrorNil	bool
	} {
		{"Empty pVal","","param",false,true},
		{"Valid pVal of 0","0","param",false,true},
		{"Valid pVal of 1","1","param",true,true},
		{"InValid pVal of 2","2","param",false,false},
		{"InValid pVal of -2","-2","param",false,false},
		{"InValid pVal string","t","param",false,false},
		{"InValid pVal spaces","               ","param",false,false},
		{"InValid pVal non-ascii","网络","param",false,false},
		{"Valid pVal with empty pName","1","",true,true},
	}

	for _,test := range tests {
		testutils.OutputTestNote( t, test.testDesc )
		result,err := paramvalidation.BoolParam( test.pVal, test.pName )
		assert.Equal( t, test.expected, result )
		if test.isErrorNil {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}

}

func TestParamValidation_Uint64Param( t *testing.T ) {

		var tests = []struct{
		testDesc	string
		pVal  		string
		pName		string
		allowZero	bool
		defaultVal	uint64
		expected  	uint64
		isErrorNil	bool
	} {
		{"Empty pVal","","param",false,100,100,true},
		{"Valid pVal of 0","0","param",true,100,0,true},
		{"Valid pVal of 1","1","param",false,100,1,true},
		{"InValid pVal of 2","2","param",false,100,2,true},
		{"InValid pVal of -2","-2","param",false,100,0,false},
		{"InValid pVal string","t","param",false,100,0,false},
		{"InValid pVal spaces","               ","param",false,100,0,false},
		{"InValid pVal non-ascii","网络","param",false,100,0,false},
		{"Valid pVal with empty pName","1","",false,100,1,true},
		{"Valid pVal of 0 but AllowZero False","0","param",false,100,0,false},
	}

	for _,test := range tests {
		testutils.OutputTestNote( t, test.testDesc )
		result,err := paramvalidation.Uint64Param( test.pVal, test.pName, test.allowZero, test.defaultVal )
		assert.Equal( t, test.expected, result )
		if test.isErrorNil {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}

}

func TestParamValidation_StringParam( t *testing.T ) {

	var tests = []struct{
	testDesc	string
	pVal  		string
	pName		string
	defaultVal	string
	options 	[]string
	expected  	string
} {
	{"Empty pVal","","param","",[]string{},""},
	{"Spaces","   ","param","",[]string{},""},
	{"Valid String","abc","param","",[]string{},"abc"},
	{"Invalid pVal with Default","","param","abc",[]string{},"abc"},
	{"Non-Ascii","网络","param","",[]string{},"网络"},
	{"Numerical value","1","param","",[]string{},"1"},
	{"Valid String not in options","abc","param","default",[]string{ "val1", "val2" },"default"},
	{"Valid String in options","val2","param","default",[]string{ "val1", "val2" },"val2"},
	{"Empty String not in options","","param","default",[]string{ "val1", "val2" },"default"},
}

for _,test := range tests {
	testutils.OutputTestNote( t, test.testDesc )
	result := paramvalidation.StringParam( test.pVal, test.pName, test.defaultVal, test.options )
	assert.Equal( t, test.expected, result )
}

}