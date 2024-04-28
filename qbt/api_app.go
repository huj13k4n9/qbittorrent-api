package qbt

import (
	"encoding/json"
	wrapper "github.com/pkg/errors"
)

const (
	VersionEndpoint         = "app/version"
	WebAPIVersionEndpoint   = "app/webapiVersion"
	BuildInfoEndpoint       = "app/buildInfo"
	ShutdownEndpoint        = "app/shutdown"
	GetPreferencesEndpoint  = "app/preferences"
	SetPreferencesEndpoint  = "app/setPreferences"
	DefaultSavePathEndpoint = "app/defaultSavePath"
)

// Version get qBitTorrent application version
//
// Example: v4.6.4
func (client *Client) Version() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(VersionEndpoint, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get qbittorrent version failed")
	}

	return string(resp), nil
}

// APIVersion get qBitTorrent API version
//
// Example: 2.9.3
func (client *Client) APIVersion() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(WebAPIVersionEndpoint, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get WebAPI version failed")
	}

	return string(resp), nil
}

// GetBuildInfo get qBitTorrent build info
func (client *Client) GetBuildInfo() (BuildInfo, error) {
	if !client.Authenticated {
		return BuildInfo{}, ErrUnauthenticated
	}

	resp, err := client.Get(BuildInfoEndpoint, nil)
	if err != nil {
		return BuildInfo{}, wrapper.Wrap(err, "get build info failed")
	}

	data := BuildInfo{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return BuildInfo{}, err
	}

	return data, nil
}

// Shutdown turn qBitTorrent off
func (client *Client) Shutdown() error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(ShutdownEndpoint, nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "shutdown failed")
	}

	client.Authenticated = false
	return nil
}

// GetPreferences get qBitTorrent preferences
func (client *Client) GetPreferences() (Preferences, error) {
	if !client.Authenticated {
		return Preferences{}, ErrUnauthenticated
	}

	resp, err := client.Get(GetPreferencesEndpoint, nil)
	if err != nil {
		return Preferences{}, wrapper.Wrap(err, "get preferences failed")
	}

	data := Preferences{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Preferences{}, err
	}

	return data, nil
}

// SetPreferences set qBitTorrent preferences
func (client *Client) SetPreferences(data Preferences) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := client.Post(
		SetPreferencesEndpoint,
		map[string]string{"json": string(bytes)},
		nil,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "set preferences failed")
	}

	return nil
}

// DefaultSavePath get default save path of downloaded content
func (client *Client) DefaultSavePath() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(DefaultSavePathEndpoint, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get default save path failed")
	}

	return string(resp), nil
}
