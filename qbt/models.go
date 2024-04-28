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
