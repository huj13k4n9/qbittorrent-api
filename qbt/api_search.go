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
func (client *Client) StartSearch(pattern string, plugins []string, category string) (int, error) {
	pluginString := strings.Join(plugins, "|")

	resp, err := client.RequestAndHandleError(
		"POST", consts.StartSearchEndpoint, map[string]string{
			"pattern":  pattern,
			"plugins":  pluginString,
			"category": category,
		}, nil,
		map[string]string{
			"409":  "user has reached the limit of max running searches",
			"!200": "start search failed",
		})

	if err != nil {
		return 0, err
	}

	var result struct {
		ID int `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}

// StopSearch is used to stop a search task that is
// started using StartSearch in qBittorrent search
// engine.
//
// Use GetSearchResults with result ID to get search
// result.
func (client *Client) StopSearch(id int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.StopSearchEndpoint, map[string]string{
			"id": strconv.Itoa(id),
		}, nil,
		map[string]string{
			"404":  "search job was not found",
			"!200": "stop search failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// GetSearchStatus is used to obtain the status of search
// tasks in qBittorrent search engine.
//
// Argument `id` is optional. `id` will be ignored in
// request when it's set to 0, and the server should return
// status of all search tasks.
func (client *Client) GetSearchStatus(id int) ([]SearchStatus, error) {
	params := make(map[string]string)
	if id != 0 {
		params["id"] = strconv.Itoa(id)
	}

	resp, err := client.RequestAndHandleError(
		"POST", consts.SearchStatusEndpoint, params, nil,
		map[string]string{
			"404":  "search job was not found",
			"!200": "get search status failed",
		})

	if err != nil {
		return nil, err
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
func (client *Client) GetSearchResults(id int, limit int, offset int) (SearchResponse, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if offset != 0 {
		params["offset"] = strconv.Itoa(offset)
	}

	resp, err := client.RequestAndHandleError(
		"POST", consts.SearchResultsEndpoint, params, nil,
		map[string]string{
			"404":  "search job was not found",
			"409":  "offset is too large, or too small",
			"!200": "get search results failed",
		})

	if err != nil {
		return SearchResponse{}, err
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
func (client *Client) DeleteSearch(id int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.DeleteSearchEndpoint, map[string]string{
			"id": strconv.Itoa(id),
		}, nil,
		map[string]string{
			"404":  "search job was not found",
			"!200": "delete search failed",
		})

	if err != nil {
		return err
	}

	return nil
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
	sourceString := strings.Join(sources, "|")

	_, err := client.RequestAndHandleError(
		"POST", consts.InstallSearchPluginEndpoint, map[string]string{
			"sources": sourceString,
		}, nil,
		map[string]string{"!200": "install search plugins failed"})

	if err != nil {
		return err
	}

	return nil
}

// UninstallPlugins takes a slice of plugin names and uninstalls
// them from qBittorrent search engine.
func (client *Client) UninstallPlugins(names []string) error {
	namesString := strings.Join(names, "|")

	_, err := client.RequestAndHandleError(
		"POST", consts.UninstallSearchPluginEndpoint, map[string]string{
			"names": namesString,
		}, nil,
		map[string]string{"!200": "uninstall search plugins failed"})

	if err != nil {
		return err
	}

	return nil
}

// EnablePlugins takes a slice of plugin names and enable/disable
// them in qBittorrent search engine.
func (client *Client) EnablePlugins(names []string, enable bool) error {
	namesString := strings.Join(names, "|")

	_, err := client.RequestAndHandleError(
		"POST", consts.EnableSearchPluginEndpoint, map[string]string{
			"names":  namesString,
			"enable": strconv.FormatBool(enable),
		}, nil,
		map[string]string{"!200": "enable search plugins failed"})

	if err != nil {
		return err
	}

	return nil
}

// UpdatePlugins takes a slice of plugin names and update
// them if new version is available in qBittorrent search engine.
func (client *Client) UpdatePlugins() error {
	_, err := client.RequestAndHandleError(
		"POST", consts.UpdateSearchPluginEndpoint, nil, nil,
		map[string]string{"!200": "update search plugins failed"})

	if err != nil {
		return wrapper.Wrap(err, "update search plugins failed")
	}

	return nil
}
