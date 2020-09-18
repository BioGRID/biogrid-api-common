// Copyright 2020 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package paramvalidation

import (
	"errors"
	"strconv"
)

// Validate that a parameter contains only the values
// of 1 or zero, and return that as a boolean
func BoolParam( pVal, pName string ) (bool, error) {
	if len(pVal) > 0 {
		bInt, err := strconv.ParseUint( pVal, 10, 64 )
		if err == nil && (bInt == 1 || bInt == 0) {
			return bInt != 0, nil
		}
	}

	return false, errors.New( pName + ": can be only a 1 or 0" )
}

// Validate that a parameter contains only a uint
// value, and return it
func Uint64Param( pVal, pName string, allowZero bool ) (uint64, error) {
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
		}
	}

	return 0, errors.New( pName + ": must be an unsigned integer value greater than or equal to 0")
}