package entities

type FileInfo struct {
	Length      int64
	PieceLength int64
	Name        string
	Pieces      []SHAHash
}

type Torrent struct {
	Announce string
	Info     FileInfo
}
