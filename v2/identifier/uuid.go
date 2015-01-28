// Tideland Go Application Support - Identifier - UUID
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
	"crypto/md5"
	"crypto/rand"
	"fmt"

	"github.com/tideland/goas/v3/errors"
)

//--------------------
// UUID
//--------------------

// UUIDVersion represents the version of a UUID.
type UUIDVersion byte

const (
	UUIDv3 UUIDVersion = 3
	UUIDv4 UUIDVersion = 4
)

// UUID represents a universal identifier with 16 bytes.
// See http://en.wikipedia.org/wiki/Universally_unique_identifier.
type UUID [16]byte

// NewUUID returns a new UUID with based on the default version 4.
func NewUUID() UUID {
	uuid, err := NewUUIDv4()
	if err != nil {
		// Panic due to compatibility reasons.
		panic(err)
	}
	return uuid
}

// NewUUIDv3 generates a new UUID based on version 3 (MD5 hash of a namespace
// and a name).
func NewUUIDv3(ns, name []byte) (UUID, error) {
	uuid := UUID{}
	if ns == nil {
		return uuid, errors.New(ErrInvalidNamespace, errorMessages)
	}
	hash := md5.New()
	hash.Write([]byte(uuid[:]))
	hash.Write(name)
	copy(uuid[:], hash.Sum([]byte{})[:16])

	uuid.setVersion(UUIDv3)
	uuid.setVariant()
	return uuid, nil
}

// NewUUIDv4 generates a new UUID based on version 4 (strong random number).
func NewUUIDv4() (UUID, error) {
	uuid := UUID{}
	_, err := rand.Read([]byte(uuid[:]))
	if err != nil {
		return uuid, err
	}

	uuid.setVersion(UUIDv4)
	uuid.setVariant()
	return uuid, nil
}

// Version returns the version number of the UUID algorithm.
func (uuid UUID) Version() UUIDVersion {
	return UUIDVersion(uuid[6] >> 4)
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

// setVersion sets the version part of the UUID.
func (uuid UUID) setVersion(v UUIDVersion) {
	var version byte = byte(v) << 4
	uuid[6] = version | (uuid[6] & 15)
}

// setVariant sets the variant part of the UUID according to RfC 4122.
func (uuid UUID) setVariant() {
	var variant byte = 8 << 4
	uuid[8] = variant | (uuid[8] & 15)
}

// EOF
