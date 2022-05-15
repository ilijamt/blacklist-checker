package cmd

import (
	"fmt"
	"github.com/ilijamt/blacklist_checker"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available blacklists.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if err = dsnblFileExists(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, err := blacklist_checker.GetDNSBLs(conf.dnsbl)
		if err != nil {
			return err
		}
		for _, host := range hosts {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", host)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
