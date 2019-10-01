// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package validation

import (
	"strings"
	"reflect"
	"gopkg.in/go-playground/validator.v9"
)

type ValidationHandler struct {
	Validate *validator.Validate
}

// Initialize Validator
func (v *ValidationHandler) Initialize( ) {
	v.Validate = validator.New( )
	v.Validate.RegisterValidation( "notblank", NotBlank )
}

// Validate fields here and generate messages
// that can be incorporated into output later on
func (v *ValidationHandler) ValidateStruct( data interface{} ) ([]string) {

	var issues []string

	err := v.Validate.Struct(data)
	if err != nil {
		issues = v.formatValidationErrors( err )
	}  

	return issues
}

// Create easy to read field errors for failed validation
// for output to the user
func (v *ValidationHandler) formatValidationErrors( err error ) ([]string) {

	var errors []string

	for _, err := range err.(validator.ValidationErrors) {		
		errors = append( errors, v.formatValidationError( err ))	
	}

	return errors
} 

// Format each field slightly differently depending on 
// the type of error it is
func (v *ValidationHandler) formatValidationError( err validator.FieldError ) (string) {
	switch strings.ToLower(err.Tag()) {
	
	case "required" :
		return err.Field( ) + " is a required field and cannot be empty"
	
	case "ascii" :
		return err.Field( ) + " can contain only ascii characters"

	case "printascii" :
		return err.Field( ) + " can contain only printable ascii characters"

	case "email" :
		return err.Field( ) + " must be a valid email field"

	case "len" :
		return err.Field( ) + " must be of length " + err.Param( )

	case "min" :
		return err.Field( ) + " must be greater than or equal to " + err.Param( ) + " or at least " + err.Param( ) + " in length if a string"

	case "max" :
		return err.Field( ) + " must be less than or equal to " + err.Param( ) + " or at most " + err.Param( ) + " in length if a string"

	case "oneof" :
		return err.Field( ) + " must be one of the following values: " + err.Param( )

	case "alphanum" :
		return err.Field( ) + " must consist of only letters of the alphabet or numbers"

	case "notblank" :
		return err.Field( ) + " cannot be blank. That includes empty arrays and strings of only whitespace."

	case "alpha" :
		return err.Field( ) + " must contain only ascii alpha characaters."

	case "url" :
		return err.Field( ) + " must be a valid URL and must include the schema such as http:// or ftp:// or https:// etc."

	default :
		return err.Field( ) + " is not validly formatted"

	}
}

// NotBlank is the validation function for validating if the current field
// has a value or length greater than zero, or is not a space only string.
func NotBlank(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}