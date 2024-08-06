package cli

import (
	"fmt"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func AddQueryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListStats())
	cmd.AddCommand(CmdShowStats())
	// this line is used by starport scaffolding # 1

	return cmd
}
