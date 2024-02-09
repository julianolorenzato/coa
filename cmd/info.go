package cmd

import (
	"encoding/json"
	"github.com/julianolorenzato/fibit/memory"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows info about the current state",
	Run: func(cmd *cobra.Command, args []string) {
		info := memory.Info{
			Disk: struct{ Size uint32 }{
				Size: 1024,
			},
			RAM: struct {
				Size            uint32
				BytesPerAddress uint32
			}{
				Size:            4096,
				BytesPerAddress: 4,
			},
			Caches: nil,
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")
		enc.Encode(info)
	},
}
