package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LinkedRequestersKeyPrefix is the prefix to retrieve all LinkedRequesters
	LinkedRequestersKeyPrefix = "LinkedRequesters/value/"
)

// LinkedRequestersKey returns the store key to retrieve a LinkedRequesters from the index fields
func LinkedRequestersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
