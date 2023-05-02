package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DataKeyPrefix is the prefix to retrieve all Data
	DataKeyPrefix = "Data/value/"
)

// DataKey returns the store key to retrieve a Data from the index fields
func DataKey(
	address string,
	owner string,
	network string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	networkBytes := []byte(network)
	key = append(key, networkBytes...)
	key = append(key, []byte("/")...)

	return key
}
