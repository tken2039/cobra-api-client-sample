package apicallcmd

import (
	"github.com/spf13/cobra"
	apicall "github.com/tken2039/cobra-api-client-sample/apicall"
	"github.com/tken2039/cobra-api-client-sample/fake/json"
)

type apiCallOptions struct {
	jsonFilePath string
	format       string
	full         bool
}

func APICallCmd() *cobra.Command {
	opts := &apiCallOptions{}

	cmd := &cobra.Command{
		Use:   "apicall",
		Short: "This command shows the contents of the json file as an API response.",
		Run: func(cmd *cobra.Command, args []string) {
			runAPICall(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.jsonFilePath, "file", "f", "sample/sample.json", "json file path")
	cmd.Flags().StringVarP(&opts.format, "format", "", "table", "output format ('table' or 'json')")
	cmd.Flags().BoolVarP(&opts.full, "full", "", false, "full the output (default false)")

	return cmd
}

func runAPICall(opts *apiCallOptions) error {
	apiClient := json.NewAPIClientFakeJSON(opts.jsonFilePath)

	if opts.format != "json" {
		opts.format = "table"
	}

	outputOptions := apicall.OutputOptions{
		Format: opts.format,
		Full:   opts.full,
	}
	c := apicall.NewAPICall(apiClient, outputOptions)

	return c.Run()
}
