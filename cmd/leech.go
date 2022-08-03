package cmd

import (
	"errors"

	"github.com/arpitbbhayani/retorrent/leecher"
	"github.com/arpitbbhayani/retorrent/torrent"
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

		torrentData, err := torrent.ParseTorrentFile(args[0])
		if err != nil {
			return err
		}

		leecher.Leech(torrentData)
		return nil
	},
}
