package types

const (
	// ModuleName defines the module name
	ModuleName = "request"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_request"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AllowedOraclesKey      = "AllowedOracles/value/"
	AllowedOraclesCountKey = "AllowedOracles/count/"
)

const (
	AddTreasuryKey      = "AddTreasury/value/"
	AddTreasuryCountKey = "AddTreasury/count/"
)
