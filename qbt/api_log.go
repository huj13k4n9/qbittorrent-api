package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"strconv"
)

// Logs get main logs with specified log level and last known log ID.
func (client *Client) Logs(logLevel uint8, lastKnownID int) ([]MainLog, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	params := map[string]string{
		"last_known_id": strconv.Itoa(lastKnownID),
		"normal":        strconv.FormatBool(logLevel&consts.LogNormal != 0),
		"info":          strconv.FormatBool(logLevel&consts.LogInfo != 0),
		"warning":       strconv.FormatBool(logLevel&consts.LogWarning != 0),
		"critical":      strconv.FormatBool(logLevel&consts.LogCritical != 0),
	}

	resp, err := client.Get(consts.LogEndpoint, params, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get main logs failed")
	}

	var data []MainLog
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PeerLogs get peer logs with specified last known log ID.
func (client *Client) PeerLogs(lastKnownID int) ([]PeerLog, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	params := map[string]string{
		"last_known_id": strconv.Itoa(lastKnownID),
	}

	resp, err := client.Get(consts.PeerLogEndpoint, params, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get peer logs failed")
	}

	var data []PeerLog
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
