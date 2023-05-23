package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RequestRecordKeyPrefix is the prefix to retrieve all RequestRecord
	RequestRecordKeyPrefix = "RequestRecord/value/"
)

// RequestRecordKey returns the store key to retrieve a RequestRecord from the index fields
func RequestRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
