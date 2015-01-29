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
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

//--------------------
// UUID
//--------------------

const (
	UUIDv3 byte = 3
	UUIDv4 byte = 4
	UUIDv5 byte = 5

	UUIDNamespaceDNS = iota
	UUIDNamespaceURL
	UUIDNamespaceOID
	UUIDNamespaceX500
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
func NewUUIDv3(ns UUID, name []byte) (UUID, error) {
	uuid := UUID{}
	hash := md5.New()
	hash.Write(ns.dump())
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

// NewUUIDv5 generates a new UUID based on version 5 (SHA1 hash of a namespace
// and a name).
func NewUUIDv5(ns UUID, name []byte) (UUID, error) {
	uuid := UUID{}
	hash := sha1.New()
	hash.Write(ns.dump())
	hash.Write(name)
	copy(uuid[:], hash.Sum([]byte{})[:16])

	uuid.setVersion(UUIDv5)
	uuid.setVariant()
	return uuid, nil
}

// Version returns the version number of the UUID algorithm.
func (uuid UUID) Version() byte {
	return uuid[6] & 0xf0 >> 4
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

// dump creates a copy a byte slice.
func (uuid UUID) dump() []byte {
	dump := make([]byte, len(uuid))

	copy(dump, uuid[:])

	return dump
}

// String returns a hexadecimal string representation with
// standardized separators.
func (uuid UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// setVersion sets the version part of the UUID.
func (uuid *UUID) setVersion(v byte) {
	uuid[6] = (uuid[6] & 0x0f) | (v << 4)
}

// setVariant sets the variant part of the UUID according to RfC 4122.
func (uuid *UUID) setVariant() {
	uuid[8] = (uuid[8] & 0x0f) | (8 << 4)
}

// UUIDNamespace returns a namespace as UUID.
func UUIDNamespace(nsId int) UUID {
	var uuid UUID
	var ns []byte
	switch nsId {
	case UUIDNamespaceDNS:
		ns, _ = hex.DecodeString("6ba7b8109dad11d180b400c04fd430c8")
	case UUIDNamespaceURL:
		ns, _ = hex.DecodeString("6ba7b8119dad11d180b400c04fd430c8")
	case UUIDNamespaceOID:
		ns, _ = hex.DecodeString("6ba7b8129dad11d180b400c04fd430c8")
	case UUIDNamespaceX500:
		ns, _ = hex.DecodeString("6ba7b8149dad11d180b400c04fd430c8")
	default:
		panic("invalid UUID namespace identifier")
	}
	copy(uuid[:], ns)
	return uuid
}

// EOF
