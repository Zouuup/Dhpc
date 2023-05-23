package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MinerResponseKeyPrefix is the prefix to retrieve all MinerResponse
	MinerResponseKeyPrefix = "MinerResponse/value/"
)

// MinerResponseKey returns the store key to retrieve a MinerResponse from the index fields
func MinerResponseKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
