package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an artifact",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CRIEI ALGO")
	},
}
