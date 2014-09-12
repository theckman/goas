// Tideland Go Application Support - Identifier
//
// Copyright (C) 2009-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// Identifier provides different ways to produce identifiers
// like UUIDs.
//
// The UUID generation follows version 4 (based on random numbers),
// other identifier types are based on passed data or types. Here
// the individual parts are harmonized and concatenated by the
// passed seperators. It is the users responsibility to check if
// the identifier is unique in its context.
package identifier

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
	return version.New(2, 1, 1)
}

// EOF
