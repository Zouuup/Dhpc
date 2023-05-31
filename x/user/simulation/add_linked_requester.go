package simulation

import (
	"math/rand"

	"Dhpc/x/user/keeper"
	"Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddLinkedRequester(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddLinkedRequester{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddLinkedRequester simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddLinkedRequester simulation not implemented"), nil, nil
	}
}
