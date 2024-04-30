package qbt

type MainLog struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	Timestamp Time   `json:"timestamp"`
	Type      int    `json:"type"`
}

type PeerLog struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	Timestamp Time   `json:"timestamp"`
	Blocked   bool   `json:"blocked"`
	Reason    string `json:"reason"`
}
