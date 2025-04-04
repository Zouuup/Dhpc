package simulation

import (
	"math/rand"

	"github.com/DhpcChain/Dhpc/x/data/keeper"
	"github.com/DhpcChain/Dhpc/x/data/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteData(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteData{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteData simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DeleteData simulation not implemented"), nil, nil
	}
}
