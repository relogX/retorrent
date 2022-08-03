package torrent

import (
	"log"

	"github.com/arpitbbhayani/retorrent/entities"
)

func Download(t *entities.Torrent, peers []entities.Peer, dir string) error {
	log.Printf("downloading %s from %d peers\n", t.Info.Name, len(peers))
	return nil
}
