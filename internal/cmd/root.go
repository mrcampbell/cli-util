package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cli-util",
		Short: "A tool for the stuff I keep doing",
	}
)

func init() {

}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
