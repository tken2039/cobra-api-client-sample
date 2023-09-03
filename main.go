package main

import (
	"github.com/spf13/cobra"
	"github.com/tken2039/cobra-api-client-sample/cmd/apicallcmd"
)

func main() {
	rootCmd().Execute()
}

var Version string
var GoVersion string

// rootCmd is a root command for emablo.
func rootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cobrapi",
		Short: "This tool is a sample to create a pseudo API client using the Go CLI framework cobra.",
	}

	rootCmd.AddCommand(apicallcmd.APICallCmd())

	return rootCmd
}
