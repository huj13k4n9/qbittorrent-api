package qbt

import (
	"net/http"
)

type Client struct {
	http          *http.Client
	URL           string
	Authenticated bool
	Jar           http.CookieJar
}

type Peer struct {
	IP   string
	Port uint
}

type Tracker struct {
	URL           string `json:"url"`
	Status        int    `json:"status"`
	Tier          int    `json:"tier"`
	NumPeers      uint   `json:"num_peers"`
	NumSeeds      uint   `json:"num_seeds"`
	NumLeeches    uint   `json:"num_leeches"`
	NumDownloaded uint   `json:"num_downloaded"`
	Message       string `json:"msg"`
}

type BuildInfo struct {
	QTVersion         string `json:"qt"`
	LibTorrentVersion string `json:"libtorrent"`
	BoostVersion      string `json:"boost"`
	OpenSSLVersion    string `json:"openssl"`
	Bitness           uint   `json:"bitness"`
}

type MainLog struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Timestamp Time   `json:"timestamp"`
	Type      uint   `json:"type"`
}

type PeerLog struct {
	ID        uint   `json:"id"`
	IP        string `json:"ip"`
	Timestamp Time   `json:"timestamp"`
	Blocked   bool   `json:"blocked"`
	Reason    string `json:"reason"`
}

type TransferInfo struct {
	ConnectionStatus  string `json:"connection_status"`
	DHTNodes          uint   `json:"dht_nodes"`
	DownloadInfoData  uint   `json:"dl_info_data"`
	DownloadInfoSpeed uint   `json:"dl_info_speed"`
	DownloadRateLimit uint   `json:"dl_rate_limit"`
	UploadInfoData    uint   `json:"up_info_data"`
	UploadInfoSpeed   uint   `json:"up_info_speed"`
	UploadRateLimit   uint   `json:"up_rate_limit"`
}

type SearchStatus struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
	Total  uint   `json:"total"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Status  string         `json:"status"`
	Total   uint           `json:"total"`
}

type SearchResult struct {
	DescriptionLink  string `json:"descrLink"`
	FileName         string `json:"fileName"`
	FileSize         int    `json:"fileSize"`
	FileURL          string `json:"fileUrl"`
	NumberOfLeechers uint   `json:"nbLeechers"`
	NumberOfSeeders  uint   `json:"nbSeeders"`
	SiteURL          string `json:"siteUrl"`
}

type SearchPluginResult struct {
	Enabled             bool   `json:"enabled"`
	FullName            string `json:"fullName"`
	Name                string `json:"name"`
	SupportedCategories []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"supportedCategories"`
	Url     string `json:"url"`
	Version string `json:"version"`
}

type TorrentListRequest struct {
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
	AddedOn                  Time     `json:"added_on"`
	AmountLeft               uint     `json:"amount_left"`
	AutoTMM                  bool     `json:"auto_tmm"`
	Availability             float32  `json:"availability"`
	Category                 string   `json:"category"`
	Completed                uint     `json:"completed"`
	CompletionOn             Time     `json:"completion_on"`
	ContentPath              string   `json:"content_path"`
	DownloadLimit            int      `json:"dl_limit"`
	DownloadSpeed            uint     `json:"dlspeed"`
	DownloadPath             string   `json:"download_path"`
	Downloaded               uint     `json:"downloaded"`
	DownloadedSession        uint     `json:"downloaded_session"`
	ETA                      uint     `json:"eta"`
	FirstLastPiecePriority   bool     `json:"f_l_piece_prio"`
	ForceStart               bool     `json:"force_start"`
	Hash                     string   `json:"hash"`
	InactiveSeedingTimeLimit int      `json:"inactive_seeding_time_limit"`
	InfoHashV1               string   `json:"infohash_v1"`
	InfoHashV2               string   `json:"infohash_v2"`
	LastActivity             Time     `json:"last_activity"`
	MagnetURI                string   `json:"magnet_uri"`
	MaxRatio                 float64  `json:"max_ratio"`
	MaxInactiveSeedingTime   int      `json:"max_inactive_seeding_time"`
	MaxSeedingTime           int      `json:"max_seeding_time"`
	Name                     string   `json:"name"`
	NumComplete              uint     `json:"num_complete"`
	NumIncomplete            uint     `json:"num_incomplete"`
	NumLeechers              uint     `json:"num_leechs"`
	NumSeeds                 int      `json:"num_seeds"`
	Priority                 int      `json:"priority"`
	Progress                 float64  `json:"progress"`
	Ratio                    float64  `json:"ratio"`
	RatioLimit               float32  `json:"ratio_limit"`
	SavePath                 string   `json:"save_path"`
	SeedingTime              int      `json:"seeding_time"`
	SeedingTimeLimit         int      `json:"seeding_time_limit"`
	SeenComplete             Time     `json:"seen_complete"`
	SequentialDownload       bool     `json:"seq_dl"`
	Size                     uint     `json:"size"`
	State                    string   `json:"state"`
	SuperSeeding             bool     `json:"super_seeding"`
	Tags                     []string `json:"tags"`
	TimeActive               uint     `json:"time_active"`
	TotalSize                uint     `json:"total_size"`
	Tracker                  string   `json:"tracker"`
	TrackersCount            int      `json:"trackers_count"`
	UploadLimit              int      `json:"up_limit"`
	Uploaded                 uint     `json:"uploaded"`
	UploadedSession          uint     `json:"uploaded_session"`
	UploadSpeed              int      `json:"upspeed"`
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
	TimeElapsed            uint    `json:"time_elapsed"`
	TotalDownloaded        uint    `json:"total_downloaded"`
	TotalDownloadedSession uint    `json:"total_downloaded_session"`
	TotalSize              uint    `json:"total_size"`
	TotalUploaded          uint    `json:"total_uploaded"`
	TotalUploadedSession   uint    `json:"total_uploaded_session"`
	TotalWasted            uint    `json:"total_wasted"`
	UploadLimit            int     `json:"up_limit"`
	UploadSpeed            int     `json:"up_speed"`
	UploadSpeedAvg         int     `json:"up_speed_avg"`
}

type Preferences struct {
	Locale                             string                 `json:"locale"`
	CreateSubfolderEnabled             bool                   `json:"create_subfolder_enabled"`
	StartPausedEnabled                 bool                   `json:"start_paused_enabled"`
	AutoDeleteMode                     int                    `json:"auto_delete_mode"`
	PreallocateAll                     bool                   `json:"preallocate_all"`
	IncompleteFilesExt                 bool                   `json:"incomplete_files_ext"`
	AutoTMMEnabled                     bool                   `json:"auto_tmm_enabled"`
	TorrentChangedTMMEnabled           bool                   `json:"torrent_changed_tmm_enabled"`
	SavePathChangedTMMEnabled          bool                   `json:"save_path_changed_tmm_enabled"`
	CategoryChangedTMMEnabled          bool                   `json:"category_changed_tmm_enabled"`
	SavePath                           string                 `json:"save_path"`
	TempPathEnabled                    bool                   `json:"temp_path_enabled"`
	TempPath                           string                 `json:"temp_path"`
	ScanDirs                           map[string]interface{} `json:"scan_dirs"`
	ExportDir                          string                 `json:"export_dir"`
	ExportDirFin                       string                 `json:"export_dir_fin"`
	MailNotificationEnabled            bool                   `json:"mail_notification_enabled"`
	MailNotificationSender             string                 `json:"mail_notification_sender"`
	MailNotificationEmail              string                 `json:"mail_notification_email"`
	MailNotificationSMTP               string                 `json:"mail_notification_smtp"`
	MailNotificationSSLEnabled         bool                   `json:"mail_notification_ssl_enabled"`
	MailNotificationAuthEnabled        bool                   `json:"mail_notification_auth_enabled"`
	MailNotificationUsername           string                 `json:"mail_notification_username"`
	MailNotificationPassword           string                 `json:"mail_notification_password"`
	AutorunEnabled                     bool                   `json:"autorun_enabled"`
	AutorunProgram                     string                 `json:"autorun_program"`
	QueueingEnabled                    bool                   `json:"queueing_enabled"`
	MaxActiveDownloads                 int                    `json:"max_active_downloads"`
	MaxActiveTorrents                  int                    `json:"max_active_torrents"`
	MaxActiveUploads                   int                    `json:"max_active_uploads"`
	DontCountSlowTorrents              bool                   `json:"dont_count_slow_torrents"`
	SlowTorrentDownloadRateThreshold   int                    `json:"slow_torrent_dl_rate_threshold"`
	SlowTorrentUploadRateThreshold     int                    `json:"slow_torrent_ul_rate_threshold"`
	SlowTorrentInactiveTimer           int                    `json:"slow_torrent_inactive_timer"`
	MaxRatioEnabled                    bool                   `json:"max_ratio_enabled"`
	MaxRatio                           float64                `json:"max_ratio"`
	MaxRatioAct                        int                    `json:"max_ratio_act"`
	ListenPort                         int                    `json:"listen_port"`
	UPnP                               bool                   `json:"upnp"`
	RandomPort                         bool                   `json:"random_port"`
	DownloadLimit                      int                    `json:"dl_limit"`
	UploadLimit                        int                    `json:"up_limit"`
	MaxConnections                     int                    `json:"max_connec"`
	MaxConnectionsPerTorrent           int                    `json:"max_connec_per_torrent"`
	MaxUploads                         int                    `json:"max_uploads"`
	MaxUploadsPerTorrent               int                    `json:"max_uploads_per_torrent"`
	StopTrackerTimeout                 int                    `json:"stop_tracker_timeout"`
	EnablePieceExtentAffinity          bool                   `json:"enable_piece_extent_affinity"`
	BittorrentProtocol                 int                    `json:"bittorrent_protocol"`
	LimitUTPRate                       bool                   `json:"limit_utp_rate"`
	LimitTCPOverhead                   bool                   `json:"limit_tcp_overhead"`
	LimitLANPeers                      bool                   `json:"limit_lan_peers"`
	AltDownloadLimit                   int                    `json:"alt_dl_limit"`
	AltUploadLimit                     int                    `json:"alt_up_limit"`
	SchedulerEnabled                   bool                   `json:"scheduler_enabled"`
	ScheduleFromHour                   int                    `json:"schedule_from_hour"`
	ScheduleFromMin                    int                    `json:"schedule_from_min"`
	ScheduleToHour                     int                    `json:"schedule_to_hour"`
	ScheduleToMin                      int                    `json:"schedule_to_min"`
	SchedulerDays                      int                    `json:"scheduler_days"`
	DHT                                bool                   `json:"dht"`
	PeX                                bool                   `json:"pex"`
	LSD                                bool                   `json:"lsd"`
	Encryption                         int                    `json:"encryption"`
	AnonymousMode                      bool                   `json:"anonymous_mode"`
	ProxyType                          string                 `json:"proxy_type"`
	ProxyIP                            string                 `json:"proxy_ip"`
	ProxyPort                          int                    `json:"proxy_port"`
	ProxyPeerConnections               bool                   `json:"proxy_peer_connections"`
	ProxyAuthEnabled                   bool                   `json:"proxy_auth_enabled"`
	ProxyUsername                      string                 `json:"proxy_username"`
	ProxyPassword                      string                 `json:"proxy_password"`
	ProxyTorrentsOnly                  bool                   `json:"proxy_torrents_only"`
	IPFilterEnabled                    bool                   `json:"ip_filter_enabled"`
	IPFilterPath                       string                 `json:"ip_filter_path"`
	IPFilterTrackers                   bool                   `json:"ip_filter_trackers"`
	WebUIDomainList                    string                 `json:"web_ui_domain_list"`
	WebUIAddress                       string                 `json:"web_ui_address"`
	WebUIPort                          int                    `json:"web_ui_port"`
	WebUIUPnP                          bool                   `json:"web_ui_upnp"`
	WebUIUsername                      string                 `json:"web_ui_username"`
	WebUIPassword                      string                 `json:"web_ui_password,omitempty"`
	WebUICSRFProtectionEnabled         bool                   `json:"web_ui_csrf_protection_enabled"`
	WebUIClickjackingProtectionEnabled bool                   `json:"web_ui_clickjacking_protection_enabled"`
	WebUISecureCookieEnabled           bool                   `json:"web_ui_secure_cookie_enabled"`
	WebUIMaxAuthFailCount              int                    `json:"web_ui_max_auth_fail_count"`
	WebUIBanDuration                   int                    `json:"web_ui_ban_duration"`
	WebUISessionTimeout                int                    `json:"web_ui_session_timeout"`
	WebUIHostHeaderValidationEnabled   bool                   `json:"web_ui_host_header_validation_enabled"`
	BypassLocalAuth                    bool                   `json:"bypass_local_auth"`
	BypassAuthSubnetWhitelistEnabled   bool                   `json:"bypass_auth_subnet_whitelist_enabled"`
	BypassAuthSubnetWhitelist          string                 `json:"bypass_auth_subnet_whitelist"`
	AlternativeWebUIEnabled            bool                   `json:"alternative_web_ui_enabled"`
	AlternativeWebUIPath               string                 `json:"alternative_web_ui_path"`
	UseHTTPS                           bool                   `json:"use_https"`
	SSLKey                             string                 `json:"ssl_key"`
	SSLCert                            string                 `json:"ssl_cert"`
	WebUIHTTPSKeyPath                  string                 `json:"web_ui_https_key_path"`
	WebUIHTTPSCertPath                 string                 `json:"web_ui_https_cert_path"`
	DynamicDNSEnabled                  bool                   `json:"dyndns_enabled"`
	DynamicDNSService                  int                    `json:"dyndns_service"`
	DynamicDNSUsername                 string                 `json:"dyndns_username"`
	DynamicDNSPassword                 string                 `json:"dyndns_password"`
	DynamicDNSDomain                   string                 `json:"dyndns_domain"`
	RSSRefreshInterval                 int                    `json:"rss_refresh_interval"`
	RSSMaxArticlesPerFeed              int                    `json:"rss_max_articles_per_feed"`
	RSSProcessingEnabled               bool                   `json:"rss_processing_enabled"`
	RSSAutoDownloadingEnabled          bool                   `json:"rss_auto_downloading_enabled"`
	RSSDownloadRepackProperEpisodes    bool                   `json:"rss_download_repack_proper_episodes"`
	RSSSmartEpisodeFilters             string                 `json:"rss_smart_episode_filters"`
	AddTrackersEnabled                 bool                   `json:"add_trackers_enabled"`
	AddTrackers                        string                 `json:"add_trackers"`
	WebUIUseCustomHttpHeadersEnabled   bool                   `json:"web_ui_use_custom_http_headers_enabled"`
	WebUICustomHttpHeaders             string                 `json:"web_ui_custom_http_headers"`
	MaxSeedingTimeEnabled              bool                   `json:"max_seeding_time_enabled"`
	MaxSeedingTime                     int                    `json:"max_seeding_time"`
	AnnounceIP                         string                 `json:"announce_ip"`
	AnnounceToAllTiers                 bool                   `json:"announce_to_all_tiers"`
	AnnounceToAllTrackers              bool                   `json:"announce_to_all_trackers"`
	AsyncIoThreads                     int                    `json:"async_io_threads"`
	BannedIPs                          string                 `json:"banned_IPs"`
	CheckingMemoryUse                  int                    `json:"checking_memory_use"`
	CurrentInterfaceAddress            string                 `json:"current_interface_address"`
	CurrentNetworkInterface            string                 `json:"current_network_interface"`
	DiskCache                          int                    `json:"disk_cache"`
	DiskCacheTTL                       int                    `json:"disk_cache_ttl"`
	EmbeddedTrackerPort                int                    `json:"embedded_tracker_port"`
	EnableCoalesceReadWrite            bool                   `json:"enable_coalesce_read_write"`
	EnableEmbeddedTracker              bool                   `json:"enable_embedded_tracker"`
	EnableMultiConnectionsFromSameIP   bool                   `json:"enable_multi_connections_from_same_ip"`
	EnableOSCache                      bool                   `json:"enable_os_cache"`
	EnableUploadSuggestions            bool                   `json:"enable_upload_suggestions"`
	FilePoolSize                       int                    `json:"file_pool_size"`
	OutgoingPortsMax                   int                    `json:"outgoing_ports_max"`
	OutgoingPortsMin                   int                    `json:"outgoing_ports_min"`
	RecheckCompletedTorrents           bool                   `json:"recheck_completed_torrents"`
	ResolvePeerCountries               bool                   `json:"resolve_peer_countries"`
	SaveResumeDataInterval             int                    `json:"save_resume_data_interval"`
	SendBufferLowWatermark             int                    `json:"send_buffer_low_watermark"`
	SendBufferWatermark                int                    `json:"send_buffer_watermark"`
	SendBufferWatermarkFactor          int                    `json:"send_buffer_watermark_factor"`
	SocketBacklogSize                  int                    `json:"socket_backlog_size"`
	UploadChokingAlgorithm             int                    `json:"upload_choking_algorithm"`
	UploadSlotsBehavior                int                    `json:"upload_slots_behavior"`
	UPnPLeaseDuration                  int                    `json:"upnp_lease_duration"`
	UTPTCPMixedMode                    int                    `json:"utp_tcp_mixed_mode"`
}
