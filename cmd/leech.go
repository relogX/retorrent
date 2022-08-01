package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(leechCmd)
}

var leechCmd = &cobra.Command{
	Use:   "leech <torrent file>",
	Short: "run torrent client in leeching mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("torrent file is required")
		}
		torrentFile := args[0]
		log.Println("running leecher for torrent file", torrentFile)
		return nil
	},
}
