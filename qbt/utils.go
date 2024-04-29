package qbt

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Time time.Time

const URLPattern = "%s/api/v2/%s"
const Version = "v0.1"

var ErrBadResponse = errors.New("received bad response")
var ErrUnauthenticated = errors.New("unauthenticated request")

func (t *Time) UnmarshalJSON(bytes []byte) error {
	timestamp, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(timestamp, 0))
	return nil
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

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}
