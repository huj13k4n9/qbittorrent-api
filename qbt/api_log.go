package qbt

import (
	"encoding/json"
	wrapper "github.com/pkg/errors"
	"strconv"
)

const (
	LogEndpoint     = "log/main"
	PeerLogEndpoint = "log/peers"
)

// Logs get main logs with specified log level and last known log ID.
func (client *Client) Logs(logLevel uint8, lastKnownID int) ([]MainLog, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	params := map[string]string{
		"last_known_id": strconv.Itoa(lastKnownID),
		"normal":        strconv.FormatBool(logLevel&LogNormal != 0),
		"info":          strconv.FormatBool(logLevel&LogInfo != 0),
		"warning":       strconv.FormatBool(logLevel&LogWarning != 0),
		"critical":      strconv.FormatBool(logLevel&LogCritical != 0),
	}

	resp, err := client.Get(LogEndpoint, params)
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

	resp, err := client.Get(PeerLogEndpoint, params)
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
