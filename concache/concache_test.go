// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package concache_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/BioGRID/biogrid-api-common/concache"
)

func TestConcache_SetRecord( t *testing.T ) {
	var c concache.ConCache
	var emptyStruct struct{}
	c.Initialize( )
	c.Set( "test", emptyStruct )
	assert.Equal(t, c.Has("test"), true)
}

func TestConcache_GetRecord( t *testing.T ) {
	var c concache.ConCache
	type TestStruct struct {
		Test string
	}
	i := TestStruct{ Test: "input_test" }
	c.Initialize( )
	c.Set( "test", i )
	r, _ := c.Get( "test" )
	rval := r.(TestStruct).Test
	assert.Equal(t, r, i )
	assert.Equal(t, rval, "input_test" )
}

func TestConcache_RemoveRecord( t *testing.T ) {
	var c concache.ConCache
	var emptyStruct struct{}
	c.Initialize( )
	c.Set( "test", emptyStruct )
	assert.Equal(t, c.Has("test"), true)
	c.Remove( "test" )
	assert.Equal(t, c.Has("test"), false)
}

func TestConcache_CountRecords( t *testing.T ) {
	var c concache.ConCache
	var emptyStruct struct{}
	c.Initialize( )
	c.Set( "test", emptyStruct )
	assert.Equal(t, c.Count(), 1)
	c.Set( "test2", emptyStruct )
	c.Set( "test3", emptyStruct )
	assert.Equal(t, c.Count(), 3)
	c.Remove( "test2" )
	assert.Equal(t, c.Count(), 2)
}