package leecher

import (
	"log"

	"github.com/arpitbbhayani/retorrent/entities"
)

func RequestPeers(t *entities.Torrent) error {
	return nil
}

func Run(t *entities.Torrent) error {
	log.Println("leeching", t.Info.Name)
	return nil
}
