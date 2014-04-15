// Tideland Go Application Support - Scroller
//
// Copyright (C) 2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package scroller

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
	ecNoError = iota
	ecNoSource
	ecNoTarget
	ecNegativeLines

	msgNoSource      = "cannot start scroller: no source"
	msgNoTarget      = "cannot start scroller: no target"
	msgNegativeLines = "negative number of lines not allowed: %d"
)

//--------------------
// TESTING
//--------------------

// IsNoSourceError returns true, if the error signals that
// no source has been passed.
func IsNoSourceError(err error) bool {
	return errors.IsError(err, ecNoSource)
}

// IsNoTargetError returns true, if the error signals that
// no target has been passed.
func IsNoTargetError(err error) bool {
	return errors.IsError(err, ecNoTarget)
}

// IsNegativeLinesError returns true, if the error shows the
// setting of a negative number of lines to start with.
func IsNegativeLinesError(err error) bool {
	return errors.IsError(err, ecNegativeLines)
}

// EOF
