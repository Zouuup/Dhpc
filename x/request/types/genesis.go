package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AllowedOraclesList: []AllowedOracles{},
		MinerResponseList:  []MinerResponse{},
		RequestRecordList:  []RequestRecord{},
		Treasury:           nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in allowedOracles
	allowedOraclesIdMap := make(map[uint64]bool)
	allowedOraclesCount := gs.GetAllowedOraclesCount()
	for _, elem := range gs.AllowedOraclesList {
		if _, ok := allowedOraclesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for allowedOracles")
		}
		if elem.Id >= allowedOraclesCount {
			return fmt.Errorf("allowedOracles id should be lower or equal than the last id")
		}
		allowedOraclesIdMap[elem.Id] = true
	}
	// Check for duplicated index in minerResponse
	minerResponseIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinerResponseList {
		index := string(MinerResponseKey(elem.UUID))
		if _, ok := minerResponseIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for minerResponse")
		}
		minerResponseIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in requestRecord
	requestRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.RequestRecordList {
		index := string(RequestRecordKey(elem.UUID))
		if _, ok := requestRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for requestRecord")
		}
		requestRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
