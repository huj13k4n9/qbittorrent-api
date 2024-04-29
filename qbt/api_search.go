package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"strconv"
	"strings"
)

// StartSearch is used to start a search request to
// qBittorrent search engine. When this request is
// complete, an ID of result will be returned.
//
// Use GetSearchResults with result ID to get
// search result, and use StopSearch to stop a search
// task.
func (client *Client) StartSearch(pattern string, plugins []string, category string) (uint, error) {
	if !client.Authenticated {
		return 0, ErrUnauthenticated
	}

	pluginString := strings.Join(plugins, "|")

	resp, err := client.Post(consts.StartSearchEndpoint, map[string]string{
		"pattern":  pattern,
		"plugins":  pluginString,
		"category": category,
	}, nil)

	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		break
	case 409:
		return 0, wrapper.Wrap(ErrBadResponse, "user has reached the limit of max running searches")
	default:
		return 0, wrapper.Wrap(ErrBadResponse, "start search failed")
	}

	var result struct {
		ID int `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return uint(result.ID), nil
}

// StopSearch is used to stop a search task that is
// started using StartSearch in qBittorrent search
// engine.
//
// Use GetSearchResults with result ID to get search
// result.
func (client *Client) StopSearch(id uint) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(consts.StopSearchEndpoint, map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}, nil)

	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return wrapper.Wrap(ErrBadResponse, "search job was not found")
	default:
		return wrapper.Wrap(ErrBadResponse, "stop search failed")
	}
}

// GetSearchStatus is used to obtain the status of search
// tasks in qBittorrent search engine.
//
// Argument `id` is optional. `id` will be ignored in
// request when it's set to 0, and the server should return
// status of all search tasks.
func (client *Client) GetSearchStatus(id uint) ([]SearchStatus, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	params := make(map[string]string)
	if id != 0 {
		params["id"] = strconv.FormatUint(uint64(id), 10)
	}

	resp, err := client.Post(consts.SearchStatusEndpoint, params, nil)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return nil, wrapper.Wrap(ErrBadResponse, "search job was not found")
	default:
		return nil, wrapper.Wrap(ErrBadResponse, "get search status failed")
	}

	var result []SearchStatus
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSearchResults is used to obtain the results of search
// tasks in qBittorrent search engine.
//
// Arguments `limit` and `offset` are optional. `limit` will
// be ignored in request when it's set to 0 or negative (means
// no limits on results). `offset` will be ignored in request
// when it's set to 0 (means no offset in results).
func (client *Client) GetSearchResults(id uint, limit int, offset int) (SearchResponse, error) {
	if !client.Authenticated {
		return SearchResponse{}, ErrUnauthenticated
	}

	params := make(map[string]string)
	params["id"] = strconv.FormatUint(uint64(id), 10)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if offset != 0 {
		params["offset"] = strconv.Itoa(offset)
	}

	resp, err := client.Post(consts.SearchResultsEndpoint, params, nil)
	if err != nil {
		return SearchResponse{}, err
	}

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return SearchResponse{}, wrapper.Wrap(ErrBadResponse, "search job was not found")
	case 409:
		return SearchResponse{}, wrapper.Wrap(ErrBadResponse, "offset is too large, or too small")
	default:
		return SearchResponse{}, wrapper.Wrap(ErrBadResponse, "get search results failed")
	}

	var result SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return SearchResponse{}, err
	}

	return result, nil
}

// DeleteSearch is used to delete a search task created
// by StartSearch in qBittorrent search engine.
func (client *Client) DeleteSearch(id uint) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(consts.DeleteSearchEndpoint, map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}, nil)

	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return wrapper.Wrap(ErrBadResponse, "search job was not found")
	default:
		return wrapper.Wrap(ErrBadResponse, "delete search failed")
	}
}

// GetSearchPlugins retrieves a list of available search
// plugins from qBittorrent search engine.
//
// It returns a slice of SearchPluginResult, which contains
// the details of each plugin, and an error if there is any
// problem during the retrieval process.
func (client *Client) GetSearchPlugins() ([]SearchPluginResult, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	resp, err := client.Get(consts.SearchPluginsEndpoint, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "get search plugins failed")
	}

	var data []SearchPluginResult
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// InstallPlugins takes a slice of plugin source URLs and installs
// them to qBittorrent search engine. `sources` can be URL or file path.
func (client *Client) InstallPlugins(sources []string) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	sourceString := strings.Join(sources, "|")
	resp, err := client.Post(consts.InstallSearchPluginEndpoint, map[string]string{
		"sources": sourceString,
	}, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "install search plugins failed")
	}
	return nil
}

// UninstallPlugins takes a slice of plugin names and uninstalls
// them from qBittorrent search engine.
func (client *Client) UninstallPlugins(names []string) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	namesString := strings.Join(names, "|")
	resp, err := client.Post(consts.UninstallSearchPluginEndpoint, map[string]string{
		"names": namesString,
	}, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "uninstall search plugins failed")
	}
	return nil
}

// EnablePlugins takes a slice of plugin names and enable/disable
// them in qBittorrent search engine.
func (client *Client) EnablePlugins(names []string, enable bool) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	namesString := strings.Join(names, "|")
	resp, err := client.Post(consts.EnableSearchPluginEndpoint, map[string]string{
		"names":  namesString,
		"enable": strconv.FormatBool(enable),
	}, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "enable search plugins failed")
	}
	return nil
}

// UpdatePlugins takes a slice of plugin names and update
// them if new version is available in qBittorrent search engine.
func (client *Client) UpdatePlugins() error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	resp, err := client.Post(consts.UpdateSearchPluginEndpoint, nil, nil)
	if err != nil {
		return wrapper.Wrap(err, "update search plugins failed")
	}

	if resp.StatusCode != 200 {
		return wrapper.Wrap(ErrBadResponse, "update search plugins failed")
	}
	return nil
}
