package qbt

import "encoding/json"

type SyncMainData struct {
	RID               int                    `json:"rid"`
	FullUpdate        bool                   `json:"full_update"`
	Torrents          map[string]TorrentInfo `json:"torrents"`
	TorrentsRemoved   []string               `json:"torrents_removed"`
	Categories        []Category             `json:"categories"`
	CategoriesRemoved []string               `json:"categories_removed"`
	Tags              []string               `json:"tags"`
	TagsRemoved       []string               `json:"tags_removed"`
	ServerState       map[string]any         `json:"server_state"`
	// ServerState Reference: https://github.com/yosemiteyss/qbittorrent_api/blob/main/lib/src/v2/sync/dto/server_state.dart
	//ServerState       struct {
	//	AllTimeDownload      int    `json:"all_time_dl"`
	//	AllTimeUpload        int    `json:"all_time_ul"`
	//	AvgTimeQueue         int    `json:"average_time_queue"`
	//	ConnectionStatus     string `json:"connection_status"`
	//	DHTNodes             int    `json:"dht_nodes"`
	//	DownloadInfoData     int    `json:"dl_info_data"`
	//	DownloadInfoSpeed    int    `json:"dl_info_speed"`
	//	DownloadRateLimit    int    `json:"dl_rate_limit"`
	//	UploadInfoData       int    `json:"up_info_data"`
	//	UploadInfoSpeed      int    `json:"up_info_speed"`
	//	UploadRateLimit      int    `json:"up_rate_limit"`
	//	FreeSpaceOnDisk      int    `json:"free_space_on_disk"`
	//	GlobalRatio          string `json:"global_ratio"`
	//	QueuedIOJobs         int    `json:"queued_io_jobs"`
	//	Queueing             bool   `json:"queueing"`
	//	ReadCacheHits        string `json:"read_cache_hits"`
	//	ReadCacheOverload    string `json:"read_cache_overloaded"`
	//	RefreshInterval      int    `json:"refresh_interval"`
	//	TotalBuffersSize     int    `json:"total_buffers_size"`
	//	TotalPeerConnections int    `json:"total_peer_connections"`
	//	TotalQueuedSize      int    `json:"total_queued_size"`
	//	TotalWastedSessions  int    `json:"total_wasted_sessions"`
	//	UseAltSpeedLimits    bool   `json:"use_alt_speed_limits"`
	//	UseSubcategories     bool   `json:"use_subcategories"`
	//	WriteCacheOverload   string `json:"write_cache_overloaded"`
	//} `json:"server_state"`
}

type SyncPeersData struct {
	RID        int            `json:"rid"`
	FullUpdate bool           `json:"full_update"`
	ShowFlags  bool           `json:"show_flags"`
	Peers      map[string]any `json:"peers"`
	// Peers Reference: https://github.com/qbittorrent/qBittorrent/blob/master/src/webui/api/synccontroller.cpp#L725
	//Peers struct {
	//	IPPort           Peer
	//	Client           string   `json:"client"`
	//	PeerIDClient     string   `json:"peer_id_client"`
	//	Progress         float64  `json:"progress"`
	//	DownloadSpeed    int      `json:"dl_speed"`
	//	UploadSpeed      int      `json:"up_speed"`
	//	TotalDownload    int      `json:"downloaded"`
	//	TotalUpload      int      `json:"uploaded"`
	//	ConnectionType   string   `json:"connection"`
	//	Flags            string   `json:"flags"`
	//	FlagsDescription string   `json:"flags_desc"`
	//	Relevance        float64  `json:"relevance"`
	//	Files            []string `json:"files"`
	//	CountryCode      string   `json:"country_code"`
	//	CountryName      string   `json:"country"`
	//} `json:"peers"`
}

func (s *SyncMainData) UnmarshalJSON(bytes []byte) error {
	type Alias SyncMainData
	aux := &struct {
		Categories map[string]Category `json:"categories"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	if aux.Categories == nil || len(aux.Categories) == 0 {
		s.Categories = nil
	} else {
		for _, category := range aux.Categories {
			s.Categories = append(s.Categories, category)
		}
	}

	return nil
}
