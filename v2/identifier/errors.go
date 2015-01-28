// Tideland Go Application Support - Identifier - Errors
//
// Copyright (C) 2009-2015 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package identifier

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
	ErrInvalidNamespace = iota + 1
)

var errorMessages = errors.Messages{
	ErrInvalidNamespace: "invalid namespace",
}

//--------------------
// TESTING
//--------------------

// IsInvalidNamespaceError returns true, if the error signals that
// the namespace is invalid.
func IsInvalidNamespaceError(err error) bool {
	return errors.IsError(err, ErrInvalidNamespace)
}

// EOF
