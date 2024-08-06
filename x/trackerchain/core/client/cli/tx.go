package cli

import (
	"fmt"
	"github.com/Skwizi_4/TrackerChain/x/core/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcomsfwedsgmands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateStats())
	//cmd.AddCommand(CmdUpdateStats())
	//cmd.AddCommand(CmdDeleteStats())
	//cmd.AddCommand(CmdIssue())
	//cmd.AddCommand(CmdWithdraw())
	// this line is used by starport scaffolding # 1

	return cmd
}
