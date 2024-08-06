package cli

import "github.com/spf13/cobra"

func CmdShowStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats-show",
		Short: "Show statistics about a stats(id)",
	}
	return cmd
}
func CmdListStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats_list",
		Short: "List all stats",
	}
	return cmd
}
