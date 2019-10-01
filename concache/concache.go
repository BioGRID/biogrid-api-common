// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package concache

import (
	"github.com/orcaman/concurrent-map"
)

type ConCache struct {
	Store        cmap.ConcurrentMap
}

// Setup a new concurrent cache
func (c *ConCache) Initialize( ) {
	c.Store = cmap.New( )
}

// Add a new element
func (c *ConCache) Set( cacheKey string, data interface{} ) {
	c.Store.Set( cacheKey, data )
}

// Fetch a loaded element
func (c *ConCache) Get( cacheKey string ) (interface{}, bool) {
	return c.Store.Get( cacheKey )
}

// Remove an element
func (c *ConCache) Remove( cacheKey string ) {
	c.Store.Remove( cacheKey )
}

// See if a key already exists
func (c *ConCache) Has( cacheKey string ) (bool) {
	return c.Store.Has( cacheKey )
}

// See how many elements are in the map
func (c *ConCache) Count( ) (int) {
	return c.Store.Count( )
}