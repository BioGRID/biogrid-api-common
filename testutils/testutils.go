// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package testutils

import (
	"testing"
)

// Output a test note
func OutputTestNote( t *testing.T, note string ) {
	t.Logf( "TEST > %s", note )
}