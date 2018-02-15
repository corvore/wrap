package cli

import (
	"github.com/spf13/cobra"
)

var longDescription = `About:
Wrap is an open source command line tool that is able to
convert Fountain and Wrap files into a correctly formatted
screen- or stageplay as an HTML or a PDF.`

// WrapCmd represents the base command when called without any subcommands
var WrapCmd = &cobra.Command{
	Use:   "wrap",
	Short: "Generate HTML and/or PDF output from Fountain files",
	Long:  longDescription,
}

var (
	outFlag       string
	benchmarkFlag bool
)

func init() {
	// Define flags used by all subcommands
	WrapCmd.PersistentFlags().StringVarP(&outFlag, "out", "o", "", "specify the `file` name to be used")
	WrapCmd.PersistentFlags().BoolVar(&benchmarkFlag, "benchmark", false, "measure the time spend on certain tasks")
}