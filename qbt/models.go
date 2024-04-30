package qbt

import (
	"fmt"
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
	Port int
}

type Category struct {
	Name     string `json:"name"`
	SavePath string `json:"savePath"`
}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}
