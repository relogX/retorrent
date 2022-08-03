package leecher

import (
	"io/ioutil"
	"log"

	"github.com/arpitbbhayani/retorrent/entities"
	"github.com/arpitbbhayani/retorrent/torrent"
	"github.com/arpitbbhayani/retorrent/tracker"
)

// https://www.bittorrent.org/beps/bep_0003.html

func Leech(t *entities.Torrent) error {
	var err error

	log.Println("connecting to the tracker at", t.Announce)
	tr := tracker.NewTrackerClient(t)

	log.Println("finding peers to download the file")
	peers, err := tr.FetchPeers()
	if err != nil {
		return err
	}
	log.Printf("found %d peers\n", len(peers))

	dir, err := ioutil.TempDir("/tmp", "123")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("creating temporary directory %s\n", dir)

	torrent.Download(t, peers, dir)

	return nil
}
