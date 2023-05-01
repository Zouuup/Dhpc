package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SlashHistoryKeyPrefix is the prefix to retrieve all SlashHistory
	SlashHistoryKeyPrefix = "SlashHistory/value/"
)

// SlashHistoryKey returns the store key to retrieve a SlashHistory from the index fields
func SlashHistoryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
