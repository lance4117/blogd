package blog

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/lance4117/blogd/api/blog/blog"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "PostList",
					Use:            "post-list",
					Short:          "Query PostList",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "PostItem",
					Use:            "post-item",
					Short:          "Query PostItem",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "PostCount",
					Use:            "post-count",
					Short:          "Query PostCount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateBlog",
					Use:            "create-blog [title] [content]",
					Short:          "Send a CreateBlog tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "content"}},
				},
				{
					RpcMethod:      "MsgUpdateBlog",
					Use:            "msg-update-blog [id] [title] [content]",
					Short:          "Send a MsgUpdateBlog tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "title"}, {ProtoField: "content"}},
				},
				{
					RpcMethod:      "PostCount",
					Use:            "post-count",
					Short:          "Send a PostCount tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
