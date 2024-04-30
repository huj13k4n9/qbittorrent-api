package qbt

type TransferInfo struct {
	ConnectionStatus  string `json:"connection_status"`
	DHTNodes          int    `json:"dht_nodes"`
	DownloadInfoData  int    `json:"dl_info_data"`
	DownloadInfoSpeed int    `json:"dl_info_speed"`
	DownloadRateLimit int    `json:"dl_rate_limit"`
	UploadInfoData    int    `json:"up_info_data"`
	UploadInfoSpeed   int    `json:"up_info_speed"`
	UploadRateLimit   int    `json:"up_rate_limit"`
}
