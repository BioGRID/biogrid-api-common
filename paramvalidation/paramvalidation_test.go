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
		ok		bool
		isErrorNil	bool
	} {
		{"Empty pVal","","param",false,false,true},
		{"Valid pVal of 0","0","param",false,true,true},
		{"Valid pVal of 1","1","param",true,true,true},
		{"InValid pVal of 2","2","param",false,true,false},
		{"InValid pVal of -2","-2","param",false,true,false},
		{"InValid pVal string","t","param",false,true,false},
		{"InValid pVal spaces","               ","param",false,true,false},
		{"InValid pVal non-ascii","网络","param",false,true,false},
		{"Valid pVal with empty pName","1","",true,true,true},
	}

	for _,test := range tests {
		testutils.OutputTestNote( t, test.testDesc )
		result,ok,err := paramvalidation.BoolParam( test.pVal, test.pName )
		assert.Equal( t, test.expected, result )
		assert.Equal( t, test.ok, ok)
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
		expected  	uint64
		ok			bool
		isErrorNil	bool
	} {
		{"Empty pVal","","param",false,0,false,true},
		{"Valid pVal of 0","0","param",true,0,true,true},
		{"Valid pVal of 1","1","param",false,1,true,true},
		{"InValid pVal of 2","2","param",false,2,true,true},
		{"InValid pVal of -2","-2","param",false,0,true,false},
		{"InValid pVal string","t","param",false,0,true,false},
		{"InValid pVal spaces","               ","param",false,0,true,false},
		{"InValid pVal non-ascii","网络","param",false,0,true,false},
		{"Valid pVal with empty pName","1","",false,1,true,true},
		{"Valid pVal of 0 but AllowZero False","0","param",false,0,true,false},
	}

	for _,test := range tests {
		testutils.OutputTestNote( t, test.testDesc )
		result,ok,err := paramvalidation.Uint64Param( test.pVal, test.pName, test.allowZero )
		assert.Equal( t, test.expected, result )
		assert.Equal( t, test.ok, ok)
		if test.isErrorNil {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}

}