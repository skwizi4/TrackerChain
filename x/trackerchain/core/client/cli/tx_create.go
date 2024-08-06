package cli

import (
	"github.com/Skwizi_4/TrackerChain/x/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"
	"os"
)

func CmdCreateStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [path]",
		Short: "Create a new block",
		Long:  "fwegwgwgegwgewrkjgaigrehjiagiewjgikeokswirgjfjGFWEIKAGLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLWRE",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			BlockBody, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			argStats := new(types.UserResults)
			err = proto.Unmarshal(BlockBody, argStats)
			if err != nil {
				return err
			}

			clientctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientctx, cmd.Flags(), argStats)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
