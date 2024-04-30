package consts

// Authentication endpoints
const (
	LoginEndpoint  = "auth/login"
	LogoutEndpoint = "auth/logout"
)

// Application endpoints
const (
	VersionEndpoint         = "app/version"
	WebAPIVersionEndpoint   = "app/webapiVersion"
	BuildInfoEndpoint       = "app/buildInfo"
	ShutdownEndpoint        = "app/shutdown"
	GetPreferencesEndpoint  = "app/preferences"
	SetPreferencesEndpoint  = "app/setPreferences"
	DefaultSavePathEndpoint = "app/defaultSavePath"
)

// Log endpoints
const (
	LogEndpoint     = "log/main"
	PeerLogEndpoint = "log/peers"
)

// Sync endpoints
const (
	MainDataEndpoint         = "sync/main"
	TorrentPeersDataEndpoint = "sync/torrentPeers"
)

// Transfer info endpoints
const (
	GetGlobalTransferInfoEndpoint  = "transfer/info"
	GetSpeedLimitsModeEndpoint     = "transfer/speedLimitsMode"
	ToggleSpeedLimitsModeEndpoint  = "transfer/toggleSpeedLimitsMode"
	GetGlobalDownloadLimitEndpoint = "transfer/downloadLimit"
	GetGlobalUploadLimitEndpoint   = "transfer/uploadLimit"
	SetGlobalDownloadLimitEndpoint = "transfer/setDownloadLimit"
	SetGlobalUploadLimitEndpoint   = "transfer/setUploadLimit"
	BanPeersEndpoint               = "transfer/banPeers"
)

// Torrent management endpoints
const (
	GetTorrentListEndpoint            = "torrents/info"
	GetTorrentPropertiesEndpoint      = "torrents/properties"
	GetTorrentTrackersEndpoint        = "torrents/trackers"
	GetTorrentWebSeedsEndpoint        = "torrents/webseeds"
	GetTorrentContentsEndpoint        = "torrents/files"
	GetTorrentPieceStatesEndpoint     = "torrents/pieceStates"
	GetTorrentPieceHashesEndpoint     = "torrents/pieceHashes"
	PauseTorrentsEndpoint             = "torrents/pause"
	ResumeTorrentsEndpoint            = "torrents/resume"
	DeleteTorrentsEndpoint            = "torrents/delete"
	RecheckTorrentsEndpoint           = "torrents/recheck"
	ReannounceTorrentsEndpoint        = "torrents/reannounce"
	ExportTorrentEndpoint             = "torrents/export"
	AddNewTorrentEndpoint             = "torrents/add"
	AddTrackersToTorrentEndpoint      = "torrents/addTrackers"
	EditTrackersEndpoint              = "torrents/editTracker"
	RemoveTrackersEndpoint            = "torrents/removeTrackers"
	AddPeersEndpoint                  = "torrents/addPeers"
	IncreaseTorrentPriorityEndpoint   = "torrents/increasePrio"
	DecreaseTorrentPriorityEndpoint   = "torrents/decreasePrio"
	MaximalTorrentPriorityEndpoint    = "torrents/topPrio"
	MinimalTorrentPriorityEndpoint    = "torrents/bottomPrio"
	SetFilePriorityEndpoint           = "torrents/filePrio"
	GetTorrentDownloadLimitEndpoint   = "torrents/downloadLimit"
	SetTorrentDownloadLimitEndpoint   = "torrents/setDownloadLimit"
	SetTorrentShareLimitEndpoint      = "torrents/setShareLimits"
	GetTorrentUploadLimitEndpoint     = "torrents/uploadLimit"
	SetTorrentUploadLimitEndpoint     = "torrents/setUploadLimit"
	SetTorrentLocationEndpoint        = "torrents/setLocation"
	SetTorrentNameEndpoint            = "torrents/rename"
	SetTorrentCategoryEndpoint        = "torrents/setCategory"
	GetAllCategoriesEndpoint          = "torrents/categories"
	AddNewCategoryEndpoint            = "torrents/createCategory"
	EditCategoryEndpoint              = "torrents/editCategory"
	RemoveCategoriesEndpoint          = "torrents/removeCategories"
	AddTorrentTagsEndpoint            = "torrents/addTags"
	RemoveTorrentTagsEndpoint         = "torrents/removeTags"
	GetAllTagsEndpoint                = "torrents/tags"
	CreateTagsEndpoint                = "torrents/createTags"
	DeleteTagsEndpoint                = "torrents/deleteTags"
	SetAutoTorrentManagementEndpoint  = "torrents/setAutoManagement"
	ToggleSequentialDownloadEndpoint  = "torrents/toggleSequentialDownload"
	SetFirstLastPiecePriorityEndpoint = "torrents/toggleFirstLastPiecePrio"
	SetForceStartEndpoint             = "torrents/setForceStart"
	SetSuperSeedingEndpoint           = "torrents/setSuperSeeding"
	RenameFileEndpoint                = "torrents/renameFile"
	RenameFolderEndpoint              = "torrents/renameFolder"
)

// RSS endpoints
const (
	AddRSSFolderEndpoint            = "rss/addFolder"
	AddRSSFeedEndpoint              = "rss/addFeed"
	RemoveRSSItemEndpoint           = "rss/removeItem"
	MoveRSSItemEndpoint             = "rss/moveItem"
	GetAllRSSItemsEndpoint          = "rss/items"
	MarkAsReadEndpoint              = "rss/markAsRead"
	RefreshRSSItemEndpoint          = "rss/refreshItem"
	SetAutoDownloadRuleEndpoint     = "rss/setRule"
	RenameAutoDownloadRuleEndpoint  = "rss/renameRule"
	RemoveAutoDownloadRuleEndpoint  = "rss/removeRule"
	GetAllAutoDownloadRulesEndpoint = "rss/rules"
	MatchArticlesWithRuleEndpoint   = "rss/matchingArticles"
)

// Search endpoints
const (
	StartSearchEndpoint           = "search/start"
	StopSearchEndpoint            = "search/stop"
	SearchStatusEndpoint          = "search/status"
	SearchResultsEndpoint         = "search/results"
	DeleteSearchEndpoint          = "search/delete"
	SearchPluginsEndpoint         = "search/plugins"
	InstallSearchPluginEndpoint   = "search/installPlugin"
	UninstallSearchPluginEndpoint = "search/uninstallPlugin"
	EnableSearchPluginEndpoint    = "search/enablePlugin"
	UpdateSearchPluginEndpoint    = "search/updatePlugins"
)
