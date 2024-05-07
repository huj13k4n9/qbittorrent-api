package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
)

// Version get qBittorrent application version
//
// Example: v4.6.4
func (client *Client) Version() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.VersionEndpoint, nil, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get qbittorrent version failed")
	}

	return string(resp), nil
}

// APIVersion get qBittorrent API version
//
// Example: 2.9.3
func (client *Client) APIVersion() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.WebAPIVersionEndpoint, nil, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get WebAPI version failed")
	}

	return string(resp), nil
}

// GetBuildInfo get qBittorrent build info
func (client *Client) GetBuildInfo() (*BuildInfo, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	resp, err := client.Get(consts.BuildInfoEndpoint, nil, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get build info failed")
	}

	data := &BuildInfo{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Shutdown turn qBittorrent off
func (client *Client) Shutdown() error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ShutdownEndpoint, nil, nil,
		map[string]string{"!200": "shutdown failed"})

	if err != nil {
		return err
	}

	client.Authenticated = false
	return nil
}

// GetPreferences get qBittorrent preferences
func (client *Client) GetPreferences() (*Preferences, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	resp, err := client.Get(consts.GetPreferencesEndpoint, nil, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get preferences failed")
	}

	data := &Preferences{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SetPreferences set qBittorrent preferences
func (client *Client) SetPreferences(data *Preferences) error {
	bytes, err := json.Marshal(*data)
	if err != nil {
		return err
	}

	_, err = client.RequestAndHandleError(
		"POST", consts.SetPreferencesEndpoint,
		map[string]string{"json": string(bytes)}, nil,
		map[string]string{"!200": "set preferences failed"})

	if err != nil {
		return err
	}

	return nil
}

// DefaultSavePath get default save path of downloaded content
func (client *Client) DefaultSavePath() (string, error) {
	if !client.Authenticated {
		return "", ErrUnauthenticated
	}

	resp, err := client.GetResponseBody(consts.DefaultSavePathEndpoint, nil, nil)
	if err != nil {
		return "", wrapper.Wrap(err, "get default save path failed")
	}

	return string(resp), nil
}
