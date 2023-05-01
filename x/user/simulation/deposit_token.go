package simulation

import (
	"math/rand"

	"Decent/x/user/keeper"
	"Decent/x/user/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDepositToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDepositToken{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DepositToken simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DepositToken simulation not implemented"), nil, nil
	}
}
