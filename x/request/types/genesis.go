package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MinerResponseList: []MinerResponse{},
		RequestRecordList: []RequestRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in minerResponse
	minerResponseIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinerResponseList {
		index := string(MinerResponseKey(elem.Index))
		if _, ok := minerResponseIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for minerResponse")
		}
		minerResponseIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in requestRecord
	requestRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.RequestRecordList {
		index := string(RequestRecordKey(elem.Index))
		if _, ok := requestRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for requestRecord")
		}
		requestRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
