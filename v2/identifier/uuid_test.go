// Tideland Go Application Support - Identifier - Unit Tests
//
// Copyright (C) 2009-2015 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package identifier_test

//--------------------
// IMPORTS
//--------------------

import (
	"testing"

	"github.com/tideland/goas/v2/identifier"
	"github.com/tideland/gots/v3/asserts"
)

//--------------------
// TESTS
//--------------------

// Test the standard UUID.
func TestStandardUUID(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	// Asserts.
	uuid := identifier.NewUUID()
	assert.Equal(uuid.Version(), identifier.UUIDv4)
	uuidStr := uuid.String()
	assert.Equal(len(uuid), 16, "UUID length has to be 16")
	assert.Match(uuidStr, "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", "UUID has to match")
	// Check for unique creation, but only weak test.
	uuids := make(map[string]bool)
	for i := 0; i < 1000000; i++ {
		uuid = identifier.NewUUID()
		uuidStr = uuid.String()
		assert.False(uuids[uuidStr], "UUID collision must not happen")
		uuids[uuidStr] = true
	}
	// Check for copy.
	uuidA := identifier.NewUUID()
	uuidB := uuidA.Copy()
	for i := 0; i < len(uuidA); i++ {
		uuidA[i] = 0
	}
	assert.Different(uuidA, uuidB)
}

// Test UUID versions.
func TestUUIDVersions(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	ns := identifier.UUIDNamespace(identifier.UUIDNamespaceOID)
	// Asserts.
	uuidV3, err := identifier.NewUUIDv3(ns, []byte{4, 7, 1, 1})
	assert.Nil(err)
	assert.Equal(uuidV3.Version(), identifier.UUIDv3)
	uuidV4, err := identifier.NewUUIDv4()
	assert.Nil(err)
	assert.Equal(uuidV4.Version(), identifier.UUIDv4)
	uuidV5, err := identifier.NewUUIDv5(ns, []byte{4, 7, 1, 1})
	assert.Nil(err)
	assert.Equal(uuidV5.Version(), identifier.UUIDv5)
}

// EOF
