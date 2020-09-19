// Copyright 2020 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package paramvalidation

import (
	"errors"
	"strconv"
	"strings"
)

// Validate that a parameter contains only the values
// of 1 or zero, and return that as a boolean
func BoolParam( pVal, pName string ) (bool, error) {
	if len(pVal) > 0 {
		bInt, err := strconv.ParseUint( pVal, 10, 64 )
		if err == nil && (bInt == 1 || bInt == 0) {
			return bInt != 0, nil
		} else {
			return false, errors.New( pName + ": can be only a 1 or 0" )
		}
	}

	return false, nil
}

// Validate that a parameter contains only a uint
// value, and return it
func Uint64Param( pVal, pName string, allowZero bool, defaultVal uint64 ) (uint64, error) {
	if len(pVal) > 0 {
		uintVal, err := strconv.ParseUint( pVal, 10, 64 )
		if err == nil && (uintVal >= 0) {
			if allowZero {
				return uintVal, nil
			} else {
				if uintVal != 0 {
					return uintVal, nil
				} else {
					return 0, errors.New( pName + ": must be a positive integer value greater than or equal to 1" )
				}
			}
		} else {
			return 0, errors.New( pName + ": must be a positive integer value greater than or equal to 1" )
		}
	}

	return defaultVal, nil
}

// Validate that a parameter contains a string
// and return it or default value
func StringParam( pVal, pName, defaultVal string, options []string ) (string) {
	pVal = strings.TrimSpace(pVal)
	if len(pVal) > 0 {
		// If we only have a fixed set
		// of valid options
		if len(options) > 0 {
			for _, option := range options {
				if option == pVal {
					return pVal
				}
			}
		} else {
			return pVal
		}
	}

	return defaultVal
}