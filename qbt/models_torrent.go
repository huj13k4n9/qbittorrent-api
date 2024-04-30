package qbt

import (
	"encoding/json"
	"strings"
)

type Tracker struct {
	URL           string `json:"url"`
	Status        int    `json:"status"`
	Tier          int    `json:"tier"`
	NumPeers      int    `json:"num_peers"`
	NumSeeds      int    `json:"num_seeds"`
	NumLeeches    int    `json:"num_leeches"`
	NumDownloaded int    `json:"num_downloaded"`
	Message       string `json:"msg"`
}

type AddTorrentParams struct {
	TorrentURLs               []string
	TorrentFiles              []string
	SavePath                  string
	Cookie                    string
	Category                  string
	Tags                      []string
	SkipChecking              bool
	Paused                    bool
	CreateRootFolder          string
	Rename                    string
	UploadLimit               int
	DownloadLimit             int
	RatioLimit                float64
	SeedingTimeLimit          int
	AutoTMM                   bool
	SequentialDownload        bool
	FirstLastPiecePrioritized bool
	AddToTopOfQueue           bool
	StopCondition             string
	ContentLayout             string
}

type TorrentListParams struct {
	Filter          string
	WithoutCategory bool
	Category        string
	WithoutTag      bool
	Tag             string
	Sort            string
	Reverse         bool
	Limit           int
	Offset          int
	Hashes          []string
}

type TorrentInfo struct {
	AddedOn                   Time     `json:"added_on"`
	AmountLeft                int      `json:"amount_left"`
	AutoTMM                   bool     `json:"auto_tmm"`
	Availability              float32  `json:"availability"`
	Category                  string   `json:"category"`
	Completed                 int      `json:"completed"`
	CompletionOn              Time     `json:"completion_on"`
	ContentPath               string   `json:"content_path"`
	DownloadLimit             int      `json:"dl_limit"`
	DownloadSpeed             int      `json:"dlspeed"`
	DownloadPath              string   `json:"download_path"`
	Downloaded                int      `json:"downloaded"`
	DownloadedSession         int      `json:"downloaded_session"`
	ETA                       int      `json:"eta"`
	FirstLastPiecePrioritized bool     `json:"f_l_piece_prio"`
	ForceStart                bool     `json:"force_start"`
	Hash                      string   `json:"hash"`
	InactiveSeedingTimeLimit  int      `json:"inactive_seeding_time_limit"`
	InfoHashV1                string   `json:"infohash_v1"`
	InfoHashV2                string   `json:"infohash_v2"`
	LastActivity              Time     `json:"last_activity"`
	MagnetURI                 string   `json:"magnet_uri"`
	MaxRatio                  float64  `json:"max_ratio"`
	MaxInactiveSeedingTime    int      `json:"max_inactive_seeding_time"`
	MaxSeedingTime            int      `json:"max_seeding_time"`
	Name                      string   `json:"name"`
	NumComplete               int      `json:"num_complete"`
	NumIncomplete             int      `json:"num_incomplete"`
	NumLeechers               int      `json:"num_leechs"`
	NumSeeds                  int      `json:"num_seeds"`
	Priority                  int      `json:"priority"`
	Progress                  float64  `json:"progress"`
	Ratio                     float64  `json:"ratio"`
	RatioLimit                float32  `json:"ratio_limit"`
	SavePath                  string   `json:"save_path"`
	SeedingTime               int      `json:"seeding_time"`
	SeedingTimeLimit          int      `json:"seeding_time_limit"`
	SeenComplete              Time     `json:"seen_complete"`
	SequentialDownload        bool     `json:"seq_dl"`
	Size                      int      `json:"size"`
	State                     string   `json:"state"`
	SuperSeeding              bool     `json:"super_seeding"`
	Tags                      []string `json:"tags"`
	TimeActive                int      `json:"time_active"`
	TotalSize                 int      `json:"total_size"`
	Tracker                   string   `json:"tracker"`
	TrackersCount             int      `json:"trackers_count"`
	UploadLimit               int      `json:"up_limit"`
	Uploaded                  int      `json:"uploaded"`
	UploadedSession           int      `json:"uploaded_session"`
	UploadSpeed               int      `json:"upspeed"`
}

type TorrentProperties struct {
	AdditionDate           Time    `json:"addition_date"`
	Comment                string  `json:"comment"`
	CompletionDate         Time    `json:"completion_date"`
	CreatedBy              string  `json:"created_by"`
	CreationDate           Time    `json:"creation_date"`
	DownloadLimit          int     `json:"dl_limit"`
	DownloadSpeed          int     `json:"dl_speed"`
	DownloadSpeedAvg       int     `json:"dl_speed_avg"`
	DownloadPath           string  `json:"download_path"`
	ETA                    int     `json:"eta"`
	Hash                   string  `json:"hash"`
	InfoHashV1             string  `json:"infohash_v1"`
	InfoHashV2             string  `json:"infohash_v2"`
	IsPrivate              bool    `json:"is_private"`
	LastSeen               Time    `json:"last_seen"`
	Name                   string  `json:"name"`
	NumConnections         int     `json:"nb_connections"`
	NumConnectionsLimit    int     `json:"nb_connections_limit"`
	Peers                  int     `json:"peers"`
	PeersTotal             int     `json:"peers_total"`
	PieceSize              int     `json:"piece_size"`
	PiecesHave             int     `json:"pieces_have"`
	PiecesNum              int     `json:"pieces_num"`
	Reannounce             int     `json:"reannounce"`
	SavePath               string  `json:"save_path"`
	SeedingTime            int     `json:"seeding_time"`
	Seeds                  int     `json:"seeds"`
	SeedsTotal             int     `json:"seeds_total"`
	ShareRatio             float64 `json:"share_ratio"`
	TimeElapsed            int     `json:"time_elapsed"`
	TotalDownloaded        int     `json:"total_downloaded"`
	TotalDownloadedSession int     `json:"total_downloaded_session"`
	TotalSize              int     `json:"total_size"`
	TotalUploaded          int     `json:"total_uploaded"`
	TotalUploadedSession   int     `json:"total_uploaded_session"`
	TotalWasted            int     `json:"total_wasted"`
	UploadLimit            int     `json:"up_limit"`
	UploadSpeed            int     `json:"up_speed"`
	UploadSpeedAvg         int     `json:"up_speed_avg"`
}

type TorrentFileProperties struct {
	Index        int     `json:"index"`
	IsSeed       bool    `json:"is_seed"`
	Name         string  `json:"name"`
	PieceRange   []int   `json:"piece_range"`
	Priority     int     `json:"priority"`
	Progress     float64 `json:"progress"`
	Size         int     `json:"size"`
	Availability float64 `json:"availability"`
}

func (ti *TorrentInfo) UnmarshalJSON(bytes []byte) error {
	type Alias TorrentInfo

	// Define an auxiliary struct with same fields as TorrentInfo,
	// except that `Tags` is in `string` type.
	aux := &struct {
		Tags string `json:"tags"`
		*Alias
	}{
		Alias: (*Alias)(ti),
	}

	// Resolve JSON data, `Tags` is resolved in
	// `aux.Tags` instead of `aux.Alias.Tags`
	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	// Split `aux.Tags` and set result to `aux.Alias.Tags`
	if aux.Tags != "" {
		ti.Tags = strings.Split(aux.Tags, ",")
	} else {
		ti.Tags = nil
	}

	return nil
}
