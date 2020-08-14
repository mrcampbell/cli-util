package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	timestampCmd = &cobra.Command{
		Use:   "timestamp",
		Short: "a simple timestamp generator",
		// Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			t := time.Now()
			fmt.Println(t)
		},
	}
)

func init() {
	rootCmd.AddCommand(timestampCmd)
}
