// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package validation_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/BioGRID/biogrid-auth-api/testutils"
	"github.com/BioGRID/biogrid-auth-api/validation"
)

var vh validation.ValidationHandler

func init( ) {
	vh.Initialize( )
}

// No need to test these exhaustively, since the validation
// package performs its own testing, mostly just test
// that error messages are returned under certain circumstances

func TestValidate_RequiredField( t *testing.T ) {
	var v = struct {
		S string `validate:"required"`
		I int `validate:"required"`
	}{ }
	issues := vh.ValidateStruct( &v )
	assert.Equal( t, 2, len(issues))
	v.S = ""
	issues = vh.ValidateStruct( &v )
	assert.Equal( t, 2, len(issues))
	v.I = 0
	issues = vh.ValidateStruct( &v )
	assert.Equal( t, 2, len(issues))
	v.S = "a"
	issues = vh.ValidateStruct( &v )
	assert.Equal( t, 1, len(issues))
	v.I = 1
	issues = vh.ValidateStruct( &v )
	assert.Equal( t, 0, len(issues))

}

func TestValidate_AsciiField( t *testing.T ) {
	var v = struct {
		S string `validate:"ascii"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",0},
		{"newline\n",0},
		{"ｆｏｏbar",1},
		{"123abcXYZxyzABC!@#$%^&*(){}/,<>;'[]-=+-*\\~",0},
		{"ｘｙｚ０９８",1},
		{"１２３456",1},
		{"〓⽣⻓⺙⺬صچ ۓ  ݲ",1},
		{"Ãé¥ŅĂŹ",1},
		{"βϖλµ",1},
		{"   ",0},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}

}

func TestValidate_PrintableAsciiField( t *testing.T ) {
	var v = struct {
		S string `validate:"printascii"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",0},
		{"newline\n",1},
		{"ｆｏｏbar",1},
		{"123abcXYZxyzABC!@#$%^&*(){}/,<>;'[]-=+-*\\~",0},
		{"ｘｙｚ０９８",1},
		{"１２３456",1},
		{"〓⽣⻓⺙⺬صچ ۓ  ݲ",1},
		{"Ãé¥ŅĂŹ",1},
		{"βϖλµ",1},
		{"   ",0},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}

}

func TestValidate_EmailField( t *testing.T ) {
	var v = struct {
		S string `validate:"email"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",1},
		{"    ",1},
		{"a",1},
		{"test@",1},
		{"test@test",1},
		{"test@test.ca",0},
	}

	var issues []string
	for _,test := range tests {
		testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}
}

func TestValidate_LenField( t *testing.T ) {
	var v = struct {
		S string `validate:"len=5"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",1},
		{"a",1},
		{"four",1},
		{"five5",0},
		{"fives",0},
	}

	var issues []string
	for _,test := range tests {
		testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}
}

func TestValidate_OneOfField( t *testing.T ) {
	var v = struct {
		S string `validate:"oneof=a b c"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"a",0},
		{"b",0},
		{"c",0},
		{"ab",1},
		{"",1},
	}

	var issues []string
	for _,test := range tests {
		testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}
}

func TestValidate_MinField( t *testing.T ) {
	var v = struct {
		S int `validate:"min=5"`
	}{ }

	var tests = []struct{
		param  int
		expected  int
	} {
		{1,1},
		{4,1},
		{-10,1},
		{5,0},
		{10,0},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, string(test.param) )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}
}

func TestValidate_MaxField( t *testing.T ) {
	var v = struct {
		S int `validate:"max=5"`
	}{ }

	var tests = []struct{
		param  int
		expected  int
	} {
		{1,0},
		{4,0},
		{-10,0},
		{5,0},
		{6,1},
		{10,1},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, string(test.param) )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}
}

func TestValidate_AlphaNumField( t *testing.T ) {
	var v = struct {
		S string `validate:"alphanum"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",1},
		{"newline\n",1},
		{"ｆｏｏbar",1},
		{"123abcXYZxyzABC!@#$%^&*(){}/,<>;'[]-=+-*\\~",1},
		{"ｘｙｚ０９８",1},
		{"１２３456",1},
		{"〓⽣⻓⺙⺬صچ ۓ  ݲ",1},
		{"Ãé¥ŅĂŹ",1},
		{"βϖλµ",1},
		{"   ",1},
		{"agzAGZ09", 0},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}

}

func TestValidate_NotBlank( t *testing.T ) {
	var v = struct {
		S string `validate:"notblank"`
	}{ }

	var tests = []struct{
		param  string
		expected  int
	} {
		{"",1},
		{" ",1},
		{"		",1},
		{"\n",1},
		{"test",0},
	}

	var issues []string
	for _,test := range tests {
		//testutils.OutputTestNote( t, test.param )
		v.S = test.param
		issues = vh.ValidateStruct( &v )
		assert.Equal( t, test.expected, len(issues))
	}

	var va = struct {
		S []string `validate:"notblank"`
	}{}

	issues = vh.ValidateStruct( &va )
	assert.Equal( t, 1, len(issues))

	va.S = []string{ "test" }
	issues = vh.ValidateStruct( &va )
	assert.Equal( t, 0, len(issues))

}