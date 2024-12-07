package optima

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "optima/api/optima/optima"
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
					RpcMethod:      "EvaluateInvocable",
					Use:            "evaluate-invocable [invocable-name] [input-data]",
					Short:          "Query evaluateInvocable",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "invocableName"}, {ProtoField: "inputData"}},
				},

				{
					RpcMethod:      "GetJob",
					Use:            "get-job [id]",
					Short:          "Query getJob",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
