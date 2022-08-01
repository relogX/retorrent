package cmd

import (
	"errors"
	"log"

	"github.com/arpitbbhayani/retorrent/torrent"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(peekCmd)
}

var peekCmd = &cobra.Command{
	Use:   "peek <torrent file>",
	Short: "peeks into the torrent file and prints the information about it",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("torrent file is required")
		}

		torrentData, err := torrent.ParseTorrentFile(args[0])
		if err != nil {
			return err
		}

		log.Printf("name: %s\n", torrentData.Info.Name)
		log.Printf("size: %d bytes\n", torrentData.Info.Length)
		log.Printf("----------\n")
		log.Printf("announce: %s\n", torrentData.Announce)
		log.Printf("number of pieces: %d\n", len(torrentData.Info.Pieces))
		log.Printf("piece length: %d bytes\n", torrentData.Info.PieceLength)

		return nil
	},
}
