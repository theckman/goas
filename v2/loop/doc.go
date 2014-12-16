// Tideland Go Application Support - Loop
//
// Copyright (C) 2013-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// The Loop of Tideland Go Application Support is intended to
// support the developer implementing the typical Go idiom for
// concurrent applications running in a loop in the background
// and doing a select on one or more channels. Stopping those
// loops or getting aware of internal errors requires extra
// efforts. The loop package helps to control this kind of
// goroutines.
package loop

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/goas/v1/version"
)

//--------------------
// VERSION
//--------------------

// PackageVersion returns the version of the version package.
func PackageVersion() version.Version {
	return version.New(2, 1, 3)
}

// EOF
