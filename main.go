package main

import (
	"github.com/spf13/cobra"
	apiclientcmd "github.com/tken2039/cobra-api-client-sample/cmd/apicallercmd"
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

	rootCmd.AddCommand(apiclientcmd.APIClientCmd())

	return rootCmd
}
