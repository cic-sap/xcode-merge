package cmd

import (
	"github.com/cic-sap/xcode-merge/pkg"

	"github.com/spf13/cobra"
)

var project *string

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.StartMerge(*project)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	project = mergeCmd.Flags().StringP("project", "p", "", ".xcodeproj file location")
}
