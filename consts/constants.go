package consts

// Constants of log levels
const (
	LogNormal uint8 = 1 << iota
	LogInfo
	LogWarning
	LogCritical
)

// Constants of torrent state
const (
	// TorrentStateError Some error occurred, applies to paused torrents
	TorrentStateError = "error"
	// TorrentStateMissingFiles Torrent data files is missing
	TorrentStateMissingFiles = "missingFiles"
	// TorrentStateUploading Torrent is being seeded and data is being transferred
	TorrentStateUploading = "uploading"
	// TorrentStatePauseUP Torrent is paused and has finished downloading
	TorrentStatePauseUP = "pausedUP"
	// TorrentStateQueuedUP Queuing is enabled and torrent is queued for upload
	TorrentStateQueuedUP = "queuedUP"
	// TorrentStateStalledUP Torrent is being seeded, but no connection were made
	TorrentStateStalledUP = "stalledUP"
	// TorrentStateCheckingUP Torrent has finished downloading and is being checked
	TorrentStateCheckingUP = "checkingUP"
	// TorrentStateForcedUP Torrent is forced to upload and ignore queue limit
	TorrentStateForcedUP = "forcedUP"
	// TorrentStateAllocating Torrent is allocating disk space for download
	TorrentStateAllocating = "allocating"
	// TorrentStateDownloading Torrent is being downloaded and data is being transferred
	TorrentStateDownloading = "downloading"
	// TorrentStateMetaDL Torrent has just started downloading and is fetching metadata
	TorrentStateMetaDL = "metaDL"
	// TorrentStatePausedDL Torrent is paused and has NOT finished downloading
	TorrentStatePausedDL = "pausedDL"
	// TorrentStateQueuedDL Queuing is enabled and torrent is queued for download
	TorrentStateQueuedDL = "queuedDL"
	// TorrentStateStalledDL Torrent is being downloaded, but no connection were made
	TorrentStateStalledDL = "stalledDL"
	// TorrentStateCheckingDL Same as checkingUP, but torrent has NOT finished downloading
	TorrentStateCheckingDL = "checkingDL"
	// TorrentStateForcedDL Torrent is forced to download to ignore queue limit
	TorrentStateForcedDL = "forcedDL"
	// TorrentStateCheckingResumeData Checking resume data on qBt startup
	TorrentStateCheckingResumeData = "checkingResumeData"
	// TorrentStateMoving Torrent is moving to another location
	TorrentStateMoving = "moving"
	// TorrentStateUnknown Unknown status
	TorrentStateUnknown = "unknown"
)

// Constants of tracker status
const (
	// TrackerDisabled Tracker is disabled (used for DHT, PeX, and LSD)
	TrackerDisabled = iota
	// TrackerNotBeenContacted Tracker has not been contacted yet
	TrackerNotBeenContacted
	// TrackerWorking Tracker has been contacted and is working
	TrackerWorking
	// TrackerUpdating Tracker is updating
	TrackerUpdating
	// TrackerContactedButNotWorking Tracker has been contacted, but it is not working (or doesn't send proper replies)
	TrackerContactedButNotWorking
)

// Constants of torrent file priority
const (
	FilePriorityDontDownload = 0
	FilePriorityNormal       = 1
	FilePriorityHigh         = 6
	FilePriorityMax          = 7
)

// Constants of connection status of qBittorrent
const (
	ConnectionStatusConnected    = "connected"
	ConnectionStatusDisconnected = "disconnected"
	ConnectionStatusFirewalled   = "firewalled"
)

// Constants of search status of qBittorrent
const (
	SearchStatusRunning = "Running"
	SearchStatusStopped = "Stopped"
)

// Constants related to the status of alternative speed limits
const (
	AlternativeSpeedLimitsDisabled = iota
	AlternativeSpeedLimitsEnabled
)

// Constants related to items in preferences
const (
	ScanDirsToMonitoredFolder = iota
	ScanDirsToDefaultPath
)

const (
	SchedulerEveryDay = iota
	SchedulerEveryWeekday
	SchedulerEveryWeekend
	SchedulerEveryMonday
	SchedulerEveryTuesday
	SchedulerEveryWednesday
	SchedulerEveryThursday
	SchedulerEveryFriday
	SchedulerEverySaturday
	SchedulerEverySunday
)

const (
	EncryptionPreferred = iota
	EncryptionForcedOn
	EncryptionForcedOff
)

const (
	ProxyTypeDisabled = "None"
	ProxyTypeHTTP     = "HTTP"
	ProxyTypeSOCKS5   = "SOCKS5"
	ProxyTypeSOCKS4   = "SOCKS4"
)

const (
	DynamicDNSServiceDyDNS = iota
	DynamicDNSServiceNOIP
)

const (
	MaxRatioActPause = iota
	MaxRatioActRemove
)

const (
	BitTorrentProtocolTCPAndUTP = iota
	BitTorrentProtocolTCP
	BitTorrentProtocolUTP
)

const (
	UploadChokingAlgorithmRoundRobin = iota
	UploadChokingAlgorithmFastestUpload
	UploadChokingAlgorithmAntiLeech
)

const (
	UploadSlotsFixed = iota
	UploadSlotsUploadRateBased
)

const (
	UTPTCPMixedModePreferTCP = iota
	UTPTCPMixedModePeerProportional
)
