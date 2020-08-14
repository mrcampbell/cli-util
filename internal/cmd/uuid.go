package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "a simple uuid generator",
		// Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			u, _ := uuid.NewUUID()
			fmt.Println(u)
		},
	}
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}
