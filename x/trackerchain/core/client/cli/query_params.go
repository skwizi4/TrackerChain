package cli

import "github.com/spf13/cobra"

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use: "query-params",
	}
	return cmd
}
