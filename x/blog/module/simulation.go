package blog

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/lance4117/blogd/testutil/sample"
	blogsimulation "github.com/lance4117/blogd/x/blog/simulation"
	"github.com/lance4117/blogd/x/blog/types"
)

// avoid unused import issue
var (
	_ = blogsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateBlog = "op_weight_msg_create_blog"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBlog int = 100

	opWeightMsgMsgUpdateBlog = "op_weight_msg_msg_update_blog"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMsgUpdateBlog int = 100

	opWeightMsgPostCount = "op_weight_msg_post_count"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostCount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blogGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blogGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBlog int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBlog, &weightMsgCreateBlog, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBlog = defaultWeightMsgCreateBlog
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBlog,
		blogsimulation.SimulateMsgCreateBlog(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMsgUpdateBlog int
	simState.AppParams.GetOrGenerate(opWeightMsgMsgUpdateBlog, &weightMsgMsgUpdateBlog, nil,
		func(_ *rand.Rand) {
			weightMsgMsgUpdateBlog = defaultWeightMsgMsgUpdateBlog
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMsgUpdateBlog,
		blogsimulation.SimulateMsgMsgUpdateBlog(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPostCount int
	simState.AppParams.GetOrGenerate(opWeightMsgPostCount, &weightMsgPostCount, nil,
		func(_ *rand.Rand) {
			weightMsgPostCount = defaultWeightMsgPostCount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostCount,
		blogsimulation.SimulateMsgPostCount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBlog,
			defaultWeightMsgCreateBlog,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blogsimulation.SimulateMsgCreateBlog(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMsgUpdateBlog,
			defaultWeightMsgMsgUpdateBlog,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blogsimulation.SimulateMsgMsgUpdateBlog(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPostCount,
			defaultWeightMsgPostCount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blogsimulation.SimulateMsgPostCount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
