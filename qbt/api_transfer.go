package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"strconv"
	"strings"
)

// GetTransferInfo get global transfer info of qBittorrent.
func (client *Client) GetTransferInfo() (*TransferInfo, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	resp, err := client.Get(consts.GetGlobalTransferInfoEndpoint, nil, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get transfer info failed")
	}

	var data TransferInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// SpeedLimitsMode The response is 1 if alternative speed limits are enabled, 0 otherwise.
// Use ToggleSpeedLimitsMode to switch the status of alternative speed limits.
//
// When alternative speed limits is off, the values of Preferences.AltDownloadLimit and
// Preferences.AltUploadLimit should be ignored.
func (client *Client) SpeedLimitsMode() (int, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.GetSpeedLimitsModeEndpoint, nil, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get speed limits mode failed")
	}

	ret, err := strconv.ParseInt(string(resp), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

// ToggleSpeedLimitsMode will switch the status of alternative speed limits.
// Use SpeedLimitsMode to check the current status of alternative speed limits.
//
// When alternative speed limits is off, the values of Preferences.AltDownloadLimit and
// Preferences.AltUploadLimit should be ignored.
func (client *Client) ToggleSpeedLimitsMode() error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ToggleSpeedLimitsModeEndpoint, nil, nil,
		map[string]string{"!200": "toggle speed limits mode failed"})

	if err != nil {
		return wrapper.Wrap(err, "toggle speed limits mode failed")
	}

	return nil
}

// GetGlobalDownloadLimit The response is the value of current global
// download speed limit in bytes/second; this value will be zero if no
// limit is applied.
//
// When alternative speed limits is on (SpeedLimitsMode returns 1), this
// API actually returns the alternative download speed limit instead of
// global download speed limit, which is the same as Preferences.AltDownloadLimit.
//
// When alternative speed limits is off (SpeedLimitsMode returns 0), this
// API returns the global download speed limit, the same as Preferences.DownloadLimit.
func (client *Client) GetGlobalDownloadLimit() (int, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.GetGlobalDownloadLimitEndpoint, nil, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get global download limit failed")
	}

	ret, err := strconv.ParseInt(string(resp), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

// GetGlobalUploadLimit The response is the value of current global
// upload speed limit in bytes/second; this value will be zero if no
// limit is applied.
//
// When alternative speed limits is on (SpeedLimitsMode returns 1), this
// API actually returns the alternative upload speed limit instead of
// global upload speed limit, which is the same as Preferences.AltUploadLimit.
//
// When alternative speed limits is off (SpeedLimitsMode returns 0), this
// API returns the global upload speed limit, the same as Preferences.UploadLimit.
func (client *Client) GetGlobalUploadLimit() (int, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.GetGlobalUploadLimitEndpoint, nil, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get global upload limit failed")
	}

	ret, err := strconv.ParseInt(string(resp), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

// SetGlobalDownloadLimit The parameter is the value of desired global
// download speed limit in bytes/second.
//
// When alternative speed limits is on (SpeedLimitsMode returns 1), this
// API actually sets the alternative download speed limit instead of
// global download speed limit, which is the same as Preferences.AltDownloadLimit.
//
// When alternative speed limits is off (SpeedLimitsMode returns 0), this
// API sets the global download speed limit, the same as Preferences.DownloadLimit.
func (client *Client) SetGlobalDownloadLimit(limit int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetGlobalDownloadLimitEndpoint, map[string]string{
			"limit": strconv.Itoa(limit),
		}, nil,
		map[string]string{"!200": "set global download limit failed"})

	if err != nil {
		return wrapper.Wrap(err, "set global download limit failed")
	}

	return nil
}

// SetGlobalUploadLimit The parameter is the value of desired global
// upload speed limit in bytes/second.
//
// When alternative speed limits is on (SpeedLimitsMode returns 1), this
// API actually sets the alternative upload speed limit instead of
// global upload speed limit, which is the same as Preferences.AltUploadLimit.
//
// When alternative speed limits is off (SpeedLimitsMode returns 0), this
// API sets the global upload speed limit, the same as Preferences.UploadLimit.
func (client *Client) SetGlobalUploadLimit(limit int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetGlobalUploadLimitEndpoint, map[string]string{
			"limit": strconv.Itoa(limit),
		}, nil,
		map[string]string{"!200": "set global upload limit failed"})

	if err != nil {
		return wrapper.Wrap(err, "set global upload limit failed")
	}

	return nil
}

// BanPeers is used to ban the connection of specified peers.
//
// Multiple peers are separated by a pipe `|`. Each peer is a
// colon-separated `host:port`.
func (client *Client) BanPeers(peers []*Peer) error {
	var peerStrings []string

	for _, peer := range peers {
		peerStrings = append(peerStrings, peer.String())
	}

	reqParam := strings.Join(peerStrings, "|")

	_, err := client.RequestAndHandleError(
		"POST", consts.BanPeersEndpoint, map[string]string{
			"peers": reqParam,
		}, nil,
		map[string]string{"!200": "ban peers failed"})

	if err != nil {
		return wrapper.Wrap(err, "ban peers failed")
	}

	return nil
}
