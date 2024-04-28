package qbt

import (
	"encoding/json"
	wrapper "github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	GetGlobalTransferInfoEndpoint  = "transfer/info"
	GetSpeedLimitsModeEndpoint     = "transfer/speedLimitsMode"
	ToggleSpeedLimitsModeEndpoint  = "transfer/toggleSpeedLimitsMode"
	GetGlobalDownloadLimitEndpoint = "transfer/downloadLimit"
	SetGlobalDownloadLimitEndpoint = "transfer/setDownloadLimit"
	GetGlobalUploadLimitEndpoint   = "transfer/uploadLimit"
	SetGlobalUploadLimitEndpoint   = "transfer/setUploadLimit"
	BanPeersEndpoint               = "transfer/banPeers"
)

func (client *Client) GetTransferInfo() (TransferInfo, error) {
	if !client.Authenticated {
		return TransferInfo{}, ErrUnauthenticated
	}

	resp, err := client.Get(GetGlobalTransferInfoEndpoint, nil)
	if err != nil {
		return TransferInfo{}, wrapper.Wrap(err, "get transfer info failed")
	}

	var data TransferInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return TransferInfo{}, err
	}

	return data, nil
}

func (client *Client) SpeedLimitsMode() (uint, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(GetSpeedLimitsModeEndpoint, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get speed limits mode failed")
	}

	ret, err := strconv.ParseUint(string(resp), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ret), nil
}

func (client *Client) ToggleSpeedLimitsMode() error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(ToggleSpeedLimitsModeEndpoint, nil, nil)
	if err != nil {
		return wrapper.Wrap(err, "toggle speed limits mode failed")
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "toggle speed limits mode failed")
	}

	return nil
}

func (client *Client) GetGlobalDownloadLimit() (uint, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(GetGlobalDownloadLimitEndpoint, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get global download limit failed")
	}

	ret, err := strconv.ParseUint(string(resp), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ret), nil
}

func (client *Client) GetGlobalUploadLimit() (uint, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(GetGlobalUploadLimitEndpoint, nil)
	if err != nil {
		return 0, wrapper.Wrap(err, "get global upload limit failed")
	}

	ret, err := strconv.ParseUint(string(resp), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ret), nil
}

func (client *Client) SetGlobalDownloadLimit(limit uint) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(SetGlobalDownloadLimitEndpoint, map[string]string{
		"limit": strconv.Itoa(int(limit)),
	}, nil)
	if err != nil {
		return wrapper.Wrap(err, "set global download limit failed")
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "set global download limit failed")
	}

	return nil
}

func (client *Client) SetGlobalUploadLimit(limit uint) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(SetGlobalUploadLimitEndpoint, map[string]string{
		"limit": strconv.Itoa(int(limit)),
	}, nil)
	if err != nil {
		return wrapper.Wrap(err, "set global upload limit failed")
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "set global upload limit failed")
	}

	return nil
}

func (client *Client) BanPeers(peers []Peer) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	var peerStrings []string

	for _, peer := range peers {
		peerStrings = append(peerStrings, peer.String())
	}

	reqParam := strings.Join(peerStrings, "|")
	resp, err := client.Post(BanPeersEndpoint, map[string]string{
		"peers": reqParam,
	}, nil)
	if err != nil {
		return wrapper.Wrap(err, "ban peers failed")
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "ban peers failed")
	}

	return nil
}
