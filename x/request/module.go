package request

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	// this line is used by starport scaffolding # 1

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DhpcChain/Dhpc/x/request/client/cli"
	"github.com/DhpcChain/Dhpc/x/request/keeper"
	"github.com/DhpcChain/Dhpc/x/request/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// MinimumMiners defines the minimum number of miners required
const MinimumMiners = 1

// BlackWait defines the wait time in blocks
const BlockWait = 10

// depositPerRequestToken is the deposit amount per request token
const depositPerRequestToken = "1000dhpc"

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	dataKeeper    types.DataKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	dataKeeper types.DataKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		dataKeeper:     dataKeeper,
	}
}

// Deprecated: use RegisterServices
func (am AppModule) Route() sdk.Route { return sdk.Route{} }

// Deprecated: use RegisterServices
func (AppModule) QuerierRoute() string { return types.RouterKey }

// Deprecated: use RegisterServices
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(goCtx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// If the number of miners' responses where the answer is not zero is more than 2/3 of the total number of miners,
	// then change the stage to 2
	// Create a list of miners who have responded with a non-zero answer

	requestRecords, _ := am.keeper.RequestRecordAllMinerPending(ctx, &types.QueryAllRequestRecordRequest{})

	for _, requestRecord := range requestRecords.RequestRecord {
		// get all miners responses in stage 1, and check if they are valid
		if len(requestRecord.Miners) >= MinimumMiners || ctx.BlockHeight() > int64(requestRecord.GetCreatedBlock())+BlockWait {
			requestRecord.Stage = 1
			requestRecord.UpdatedBlock = ctx.BlockHeight()
			am.keeper.SetRequestRecord(ctx, requestRecord)
		}
	}

	requestRecords, _ = am.keeper.RequestRecordAllAnswerPending(ctx, &types.QueryAllRequestRecordRequest{})
	for _, requestRecord := range requestRecords.RequestRecord {
		// get all miners responses in stage 2, and check if they are valid
		var nonZeroAnswerMiners []types.MinerResponse
		var rewardedMiners []types.MinerResponse
		for _, miner := range requestRecord.Miners {
			if miner.GetAnswer() != 0 {
				sum := miner.Answer + miner.Salt
				sumStr := strconv.Itoa(int(sum))
				hasher := md5.New()
				hasher.Write([]byte(sumStr))
				answerHash := hex.EncodeToString(hasher.Sum(nil))

				if miner.GetHash() == answerHash {
					nonZeroAnswerMiners = append(nonZeroAnswerMiners, *miner)
				}
			}
		}
		if len(nonZeroAnswerMiners) > (len(requestRecord.Miners)*2/3) || ctx.BlockHeight() > int64(requestRecord.GetUpdatedBlock())+BlockWait {
			// If more than half of the miners have the same answer, then we can switch to stage 2

			// Find the most common answer
			answerCount := make(map[int32]int)
			for _, miner := range nonZeroAnswerMiners {
				answerCount[miner.Answer]++
			}
			var maxAnswerCount int
			var maxAnswer int32
			for answer, count := range answerCount {
				if count > maxAnswerCount {
					maxAnswerCount = count
					maxAnswer = answer
				}
			}

			for _, miner := range nonZeroAnswerMiners {
				if miner.Answer == maxAnswer {
					rewardedMiners = append(rewardedMiners, miner)
				}
			}

			deposit, err := sdk.ParseCoinsNormalized(depositPerRequestToken)
			if err != nil {
				log.Printf("Unable to parse coins for creating request record: %s", err)
			}

			// Get 40% of the deposit for miners
			minerAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(40).QuoRaw(100).QuoRaw(int64(len(rewardedMiners))))
			minerAmount := sdk.NewCoins(minerAmountCoin)

			// Iterate through rewardedMiners and pay each miner amount
			for _, miner := range rewardedMiners {
				minerAddress, err := sdk.AccAddressFromBech32(miner.Creator)
				if err != nil {
					log.Printf("Unable to parse miner address: %s", err)
				}

				sdkError := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minerAddress, minerAmount)
				if sdkError != nil {
					log.Printf("Unable to send coins to miner: %s", sdkError)
				}
			}

			// Iterate through rewardedMiners' dataused, find data that's used by 80% of the miners, and pay data providers for that data
			dataCounts := make(map[string]int)
			for _, miner := range rewardedMiners {
				// Dataused is an MD5 string divided by commas, split it and iterate through it
				dataUsed := strings.Split(miner.DataUsed, ",")
				for _, data := range dataUsed {
					found, _ := am.dataKeeper.GetOwnerByHash(ctx, data)
					if found {
						dataCounts[data]++
					}
				}
			}

			// Create a list of data that's used by 80% of the miners, which means the count is equal to or more than 80% of rewardedMiners
			var dataUsedBy80Percent []string
			for data, count := range dataCounts {
				if count >= len(rewardedMiners)*80/100 {
					dataUsedBy80Percent = append(dataUsedBy80Percent, data)
				}
			}

			// Get 55% of the deposit for data providers
			dataProviderAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(55).QuoRaw(100).QuoRaw(int64(len(dataUsedBy80Percent))))
			dataProviderAmount := sdk.NewCoins(dataProviderAmountCoin)

			// Iterate through dataUsedBy80Percent and pay each data provider amount

			for _, data := range dataUsedBy80Percent {
				_, owner := am.dataKeeper.GetOwnerByHash(ctx, data)
				ownerAddress, err := sdk.AccAddressFromBech32(owner)
				if err != nil {
					log.Printf("Unable to parse miner address: %s", err)
					continue
				}

				sdkError := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, dataProviderAmount)
				if sdkError != nil {
					log.Printf("Unable to send coins to miner: %s", sdkError)
					continue
				}
			}

			// TODO: Not ideal, instead we should send all remaining tokens to the treasury
			treasuryAmountCoin := sdk.NewCoin("dhpc", deposit[0].Amount.MulRaw(5).QuoRaw(100))
			treasuryAmount := sdk.NewCoins(treasuryAmountCoin)

			treasury, found := am.keeper.GetTreasury(ctx)
			if !found {
				log.Printf("Unable to find treasury")
				continue
			}

			treasuryAddress, err := sdk.AccAddressFromBech32(treasury.Address)
			if err != nil {
				log.Printf("Unable to parse treasury address: %s", err)
				continue
			}

			sdkError := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, treasuryAddress, treasuryAmount)
			if sdkError != nil {
				log.Printf("Unable to send coins to treasury: %s", sdkError)
				continue
			}

			ctx.Logger().Info("Finishing Request", "UUID", requestRecord.UUID, "miners", len(rewardedMiners), "minerAmount", minerAmount, "dataProviders", len(dataUsedBy80Percent), "dataAmount", dataProviderAmount, "treasury", treasuryAmount)

			// Switch to Stage 2
			requestRecord.Stage = 2
			requestRecord.Score = maxAnswer
			requestRecord.UpdatedBlock = ctx.BlockHeight()
			am.keeper.SetRequestRecord(ctx, requestRecord)
		}
	}

	return []abci.ValidatorUpdate{}
}
