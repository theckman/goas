// Tideland Go Application Support - Scene
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package scene

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/goas/v3/errors"
)

//--------------------
// CONSTANTS
//--------------------

const (
	ErrSceneEnded = iota + 1
	ErrTimeout
	ErrPropAlreadyExist
	ErrPropNotFound
	ErrCleanupFailed
)

var errorMessages = errors.Messages{
	ErrSceneEnded:       "scene already ended",
	ErrTimeout:          "scene %s timeout reached at %v",
	ErrPropAlreadyExist: "prop %q already exist",
	ErrPropNotFound:     "prop %q does not exist",
	ErrCleanupFailed:    "cleanup of prop %q failed",
}

//--------------------
// TESTING
//--------------------

// IsSceneEndedError returns true, if the error signals that
// the scene isn't active anymore.
func IsSceneEdedError(err error) bool {
	return errors.IsError(err, ErrSceneEnded)
}

// IsTimeoutError returns true, if the error signals that
// the scene end after an absolute timeout.
func IsTimeoutError(err error) bool {
	return errors.IsError(err, ErrTimeout)
}

// IsPropAlreadyExistError returns true, if the error signals a
// double prop key.
func IsPropAlreadyExistError(err error) bool {
	return errors.IsError(err, ErrPropAlreadyExist)
}

// IsPropNotFoundError returns true, if the error signals a
// non-existing prop.
func IsPropNotFoundError(err error) bool {
	return errors.IsError(err, ErrPropNotFound)
}

// IsCleanupFaildError returns true, if the error signals the
// failing of a prop error.
func IsCleanupFailedError(err error) bool {
	return errors.IsError(err, ErrCleanupFailed)
}

// EOF
