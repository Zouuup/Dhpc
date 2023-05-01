package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LinkedRequestersList: []LinkedRequesters{},
		SlashHistoryList:     []SlashHistory{},
		UserList:             []User{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in linkedRequesters
	linkedRequestersIndexMap := make(map[string]struct{})

	for _, elem := range gs.LinkedRequestersList {
		index := string(LinkedRequestersKey(elem.Index))
		if _, ok := linkedRequestersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for linkedRequesters")
		}
		linkedRequestersIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in slashHistory
	slashHistoryIndexMap := make(map[string]struct{})

	for _, elem := range gs.SlashHistoryList {
		index := string(SlashHistoryKey(elem.Index))
		if _, ok := slashHistoryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for slashHistory")
		}
		slashHistoryIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in user
	userIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserList {
		index := string(UserKey(elem.Account))
		if _, ok := userIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for user")
		}
		userIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
