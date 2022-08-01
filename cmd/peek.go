package cmd

import (
	"errors"
	"fmt"

	"github.com/arpitbbhayani/retorrent/torrent"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(peekCmd)
}

var peekCmd = &cobra.Command{
	Use:   "peek <torrent file>",
	Short: "peeks into the torrent fike and prints the stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("torrent file is required")
		}

		torrentData, err := torrent.ParseTorrentFile(args[0])
		if err != nil {
			return err
		}

		fmt.Println(torrentData.Info.Length)
		fmt.Println(torrentData.Info.PieceLength)
		fmt.Println(torrentData.Info.Length / torrentData.Info.PieceLength)
		fmt.Println(len(torrentData.Info.Pieces))
		return nil
	},
}
