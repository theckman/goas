// Tideland Go Application Support - Identifier
//
// Copyright (C) 2009-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package identifier

//--------------------
// IMPORTS
//--------------------

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unicode"

	"github.com/tideland/goas/v1/version"
)

//--------------------
// VERSION
//--------------------

// PackageVersion returns the version of the version package.
func PackageVersion() version.Version {
	return version.New(2, 1, 0)
}

//--------------------
// UUID
//--------------------

// UUID represent a universal identifier with 16 bytes.
type UUID [16]byte

// NewUUID generates a new UUID based on version 4 (strong random number).
// See http://en.wikipedia.org/wiki/Universally_unique_identifier.
func NewUUID() UUID {
	uuid := UUID{}
	if _, err := io.ReadFull(rand.Reader, []byte(uuid[0:16])); err != nil {
		panic(err)
	}
	// Set version (4) and variant (2) according to RfC 4122.
	var version byte = 4 << 4
	var variant byte = 8 << 4
	uuid[6] = version | (uuid[6] & 15)
	uuid[8] = variant | (uuid[8] & 15)
	return uuid
}

// Copy returns a copy of the UUID.
func (uuid UUID) Copy() UUID {
	uuidCopy := uuid
	return uuidCopy
}

// Raw returns a copy of the UUID bytes.
func (uuid UUID) Raw() [16]byte {
	return [16]byte(uuid)
}

// String returns a hexadecimal string representation with
// standardized separators.
func (uuid UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

//--------------------
// MORE ID FUNCTIONS
//--------------------

// LimitedSepIdentifier builds an identifier out of multiple parts,
// all as lowercase strings and concatenated with the separator
// Non letters and digits are exchanged with dashes and
// reduced to a maximum of one each. If limit is true only
// 'a' to 'z' and '0' to '9' are allowed.
func LimitedSepIdentifier(sep string, limit bool, parts ...interface{}) string {
	iparts := make([]string, 0)
	for _, p := range parts {
		tmp := strings.Map(func(r rune) rune {
			// Check letter and digit.
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				lcr := unicode.ToLower(r)
				if limit {
					// Only 'a' to 'z' and '0' to '9'.
					if lcr <= unicode.MaxASCII {
						return lcr
					} else {
						return ' '
					}
				} else {
					// Every char is allowed.
					return lcr
				}
			}
			return ' '
		}, fmt.Sprintf("%v", p))
		// Only use non-empty identifier parts.
		if ipart := strings.Join(strings.Fields(tmp), "-"); len(ipart) > 0 {
			iparts = append(iparts, ipart)
		}
	}
	return strings.Join(iparts, sep)
}

// SepIdentifier builds an identifier out of multiple parts, all
// as lowercase strings and concatenated with the separator
// Non letters and digits are exchanged with dashes and
// reduced to a maximum of one each.
func SepIdentifier(sep string, parts ...interface{}) string {
	return LimitedSepIdentifier(sep, false, parts...)
}

// Identifier works like SepIdentifier but the seperator
// is set to be a colon.
func Identifier(parts ...interface{}) string {
	return SepIdentifier(":", parts...)
}

// JoinedIdentifier builds a new identifier, joinded with the
// colon as the seperator.
func JoinedIdentifier(identifiers ...string) string {
	return strings.Join(identifiers, ":")
}

// TypeAsIdentifierPart transforms the name of the arguments type into
// a part for identifiers. It's splitted at each uppercase char,
// concatenated with dashes and transferred to lowercase.
func TypeAsIdentifierPart(i interface{}) string {
	var buf bytes.Buffer
	fullTypeName := reflect.TypeOf(i).String()
	lastDot := strings.LastIndex(fullTypeName, ".")
	typeName := fullTypeName[lastDot+1:]
	for i, r := range typeName {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('-')
			}
		}
		buf.WriteRune(r)
	}
	return strings.ToLower(buf.String())
}

// EOF
