package qbt

// Constants of log levels
const (
	LogNormal uint8 = 1 << iota
	LogInfo
	LogWarning
	LogCritical
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