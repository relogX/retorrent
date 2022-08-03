package tracker

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	bencode1 "github.com/arpitbbhayani/retorrent/bencode"
	"github.com/arpitbbhayani/retorrent/entities"
	"github.com/jackpal/bencode-go"
)

type TrackerClient struct {
	t *entities.Torrent
}

func NewTrackerClient(t *entities.Torrent) *TrackerClient {
	return &TrackerClient{t}
}

func safeInt64(v interface{}) int64 {
	x, ok := v.(int64)
	if !ok {
		return 0
	}
	return x
}

func safeString(v interface{}) string {
	x, ok := v.(string)
	if !ok {
		return ""
	}
	return x
}

func prepareTrackerResponse(body io.ReadCloser) (*entities.TrackerResponse, error) {
	data, err := bencode1.Decode(bufio.NewReader(body))
	if err != nil {
		return nil, err
	}

	responseMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("response not a valid bencoding")
	}

	var peers []entities.Peer

	peersRaw, ok := responseMap["peers"].([]interface{})
	if ok {
		for _, pInterface := range peersRaw {
			pRaw, ok := pInterface.(map[string]interface{})
			if !ok {
				continue
			}
			peers = append(peers, entities.Peer{
				IP:     safeString(pRaw["ip"]),
				Port:   safeInt64(pRaw["port"]),
				PeerID: safeString(pRaw["peer id"]),
			})
		}
	}

	response := &entities.TrackerResponse{
		Complete:      safeInt64(responseMap["complete"]),
		Incomplete:    safeInt64(responseMap["incomplete"]),
		Interval:      safeInt64(responseMap["interval"]),
		FailureReason: safeString(responseMap["interval"]),
		Peers:         peers,
	}

	return response, nil
}

func (tr *TrackerClient) FetchPeers() ([]entities.Peer, error) {
	base, err := url.Parse(tr.t.Announce)
	if err != nil {
		return nil, nil
	}

	// TODO: Write encoder of bencoding instead of using Jackpal library
	var infoBuffer bytes.Buffer
	err = bencode.Marshal(&infoBuffer, tr.t.InfoRaw)
	if err != nil {
		return nil, err
	}
	infoHash := sha1.Sum(infoBuffer.Bytes())

	// request params
	params := url.Values{
		"info_hash":  []string{string(infoHash[:])},
		"peer_id":    []string{string("~RelogX Toy Torrent~")},
		"port":       []string{strconv.Itoa(6881)},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"left":       []string{strconv.Itoa(int(tr.t.Info.Length))},
	}
	base.RawQuery = params.Encode()
	url := base.String()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tResponse, err := prepareTrackerResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	if tResponse.FailureReason != "" {
		return nil, errors.New(tResponse.FailureReason)
	}

	return tResponse.Peers, nil
}
