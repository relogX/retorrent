package torrent

import (
	"bufio"
	"errors"
	"os"

	"github.com/arpitbbhayani/retorrent/bencode"
	"github.com/arpitbbhayani/retorrent/entities"
)

func batch(data []byte, batch int) []entities.SHAHash {
	var result []entities.SHAHash
	for i := 0; i < len(data); i += batch {
		hash := entities.SHAHash{}
		end := i + batch
		if end > len(data) {
			end = len(data)
		}
		copy(hash[:], data[i:end])
		result = append(result, hash)
	}
	return result
}

func getTorrent(bReader *bufio.Reader) (*entities.Torrent, error) {
	data, err := bencode.Decode(bReader)
	if err != nil {
		return nil, err
	}

	tData, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid torrent file")
	}

	tInfoData, ok := tData["info"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid torrent file")
	}

	torrent := &entities.Torrent{}
	torrent.Announce = tData["announce"].(string)
	torrent.Info = entities.FileInfo{
		PieceLength: tInfoData["piece length"].(int64),
		Length:      tInfoData["length"].(int64),
		Name:        tInfoData["name"].(string),
		Pieces:      batch([]byte(tInfoData["pieces"].(string)), 20),
	}
	torrent.InfoRaw = tData["info"].(map[string]interface{})

	return torrent, nil
}

func ParseTorrentFile(fpath string) (*entities.Torrent, error) {
	fp, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return getTorrent(bufio.NewReader(fp))
}
