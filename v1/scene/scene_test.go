// Tideland Go Application Support - Scene - Unit Tests
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package scene_test

//--------------------
// IMPORTS
//--------------------

import (
	"errors"
	"testing"

	"github.com/tideland/goas/v1/scene"
	"github.com/tideland/gots/v3/asserts"
)

//--------------------
// TESTS
//--------------------

// TestSimpleNoTimeout tests a simple scene usage without
// any timeout.
func TestSimpleNoTimeout(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, false)
	scn := scene.StartScene()

	err := scn.Store("foo", 4711)
	assert.Nil(err)
	foo, err := scn.Fetch("foo")
	assert.Nil(err)
	assert.Equal(foo, 4711)
	_, err = scn.Fetch("bar")
	assert.True(scene.IsPropNotFoundError(err))
	err = scn.Store("foo", "bar")
	assert.True(scene.IsPropAlreadyExistError(err))
	_, err = scn.Dispose("bar")
	assert.True(scene.IsPropNotFoundError(err))
	foo, err = scn.Dispose("foo")
	assert.Nil(err)
	assert.Equal(foo, 4711)
	_, err = scn.Fetch("foo")
	assert.True(scene.IsPropNotFoundError(err))

	err = scn.Stop()
	assert.Nil(err)
}

// TestCleanupNoError tests the cleanup of props with
// no errors.
func TestCleanupNoError(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, false)
	cleanups := make(map[string]interface{})
	cleanup := func(key string, prop interface{}) error {
		cleanups[key] = prop
		return nil
	}
	scn := scene.StartScene()

	err := scn.StoreClean("foo", 4711, cleanup)
	assert.Nil(err)
	err = scn.StoreClean("bar", "yadda", cleanup)
	assert.Nil(err)

	foo, err := scn.Dispose("foo")
	assert.Nil(err)
	assert.Equal(foo, 4711)

	err = scn.Stop()
	assert.Nil(err)

	assert.Length(cleanups, 2)
	assert.Equal(cleanups["foo"], 4711)
	assert.Equal(cleanups["bar"], "yadda")
}

// TestCleanupWithErrors tests the cleanup of props with errors.
func TestCleanupWithErrors(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, false)
	cleanup := func(key string, prop interface{}) error {
		return errors.New("ouch")
	}
	scn := scene.StartScene()

	err := scn.StoreClean("foo", 4711, cleanup)
	assert.Nil(err)
	err = scn.StoreClean("bar", true, cleanup)
	assert.Nil(err)
	err = scn.StoreClean("yadda", "OK", cleanup)
	assert.Nil(err)

	foo, err := scn.Dispose("foo")
	assert.True(scene.IsCleanupFailedError(err))
	assert.Nil(foo)
	bar, err := scn.Dispose("bar")
	assert.True(scene.IsCleanupFailedError(err))
	assert.Nil(bar)

	err = scn.Stop()
	assert.NotNil(err)
	assert.True(scene.IsCleanupFailedError(err))
}

// EOF
