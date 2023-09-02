package apiclientcmd

import (
	"github.com/spf13/cobra"
	"github.com/tken2039/cobra-api-client-sample/apicaller"
	"github.com/tken2039/cobra-api-client-sample/fake/json"
)

type apiClientOptions struct {
	jsonFilePath string
	format       string
	full         bool
}

func APIClientCmd() *cobra.Command {
	opts := &apiClientOptions{}

	cmd := &cobra.Command{
		Use:   "apicaller",
		Short: "This command shows the contents of the json file as an API response.",
		Run: func(cmd *cobra.Command, args []string) {
			runAPICaller(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.jsonFilePath, "file", "f", "sample/sample.json", "json file path (required))")
	cmd.MarkPersistentFlagRequired("file")
	cmd.Flags().StringVarP(&opts.format, "format", "", "table", "output format ('table' or 'json')")
	cmd.Flags().BoolVarP(&opts.full, "full", "", false, "full the output (default false)")

	return cmd
}

func runAPICaller(opts *apiClientOptions) error {
	apiClient := json.NewAPIClientFakeJSON(opts.jsonFilePath)

	if opts.format != "json" {
		opts.format = "table"
	}

	outputOptions := apicaller.OutputOptions{
		Format: opts.format,
		Full:   opts.full,
	}
	c := apicaller.NewAPICaller(apiClient, outputOptions)

	return c.Run()
}
