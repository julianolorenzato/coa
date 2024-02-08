package cmd

import "github.com/spf13/cobra"

func Execute() {
	rootCmd = &cobra.Command{
		Use: "csimu",
		Short: "",
	}
}