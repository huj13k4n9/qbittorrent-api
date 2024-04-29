package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	"strconv"
	"strings"
)

// BuildTorrentListQuery is used to check request parameters
// of torrent list API, and then turn actual values into
// `map[string]string` for Client.Get to use.
func BuildTorrentListQuery(req *TorrentListRequest) map[string]string {
	ret := make(map[string]string)

	if req.Filter != "" {
		ret["filter"] = req.Filter
	}

	if req.WithoutCategory {
		ret["category"] = ""
	} else {
		if req.Category != "" {
			ret["category"] = req.Category
		}
	}

	if req.WithoutTag {
		ret["tag"] = ""
	} else {
		if req.Tag != "" {
			ret["tag"] = req.Tag
		}
	}

	if req.Sort != "" {
		ret["sort"] = req.Sort
	}

	ret["reverse"] = strconv.FormatBool(req.Reverse)

	if req.Limit > 0 {
		ret["limit"] = strconv.Itoa(req.Limit)
	}

	if req.Offset != 0 {
		ret["offset"] = strconv.Itoa(req.Offset)
	}

	if req.Hashes != nil && len(req.Hashes) != 0 {
		ret["hashes"] = strings.Join(req.Hashes, "|")
	}

	return ret
}

// Torrents method is used to get torrent list in qBittorrent.
// Return basic info of all torrents in TorrentInfo struct.
//
// If you want to request directly with no parameters, pass
// options with nil. Otherwise, construct your own TorrentListRequest
// data is needed.
func (client *Client) Torrents(options *TorrentListRequest) ([]TorrentInfo, error) {
	var params map[string]string
	if options == nil {
		params = BuildTorrentListQuery(&TorrentListRequest{
			Filter:          "",
			WithoutCategory: false,
			WithoutTag:      false,
			Category:        "",
			Tag:             "",
			Sort:            "",
			Reverse:         false,
			Limit:           0,
			Offset:          0,
			Hashes:          nil,
		})
	} else {
		params = BuildTorrentListQuery(options)
	}

	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentListEndpoint, params, nil,
		map[string]string{"!200": "get torrents list failed"})

	if err != nil {
		return nil, err
	}

	var data []TorrentInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TorrentProperties method is used to get generic properties of specified torrent.
// Pass hash of torrent as parameter so that server can identify the torrent.
//
// Note: -1 is returned if the type of the property is integer but its value is not known.
func (client *Client) TorrentProperties(hash string) (TorrentProperties, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentPropertiesEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrents properties failed"})

	if err != nil {
		return TorrentProperties{}, err
	}

	var data TorrentProperties
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return TorrentProperties{}, err
	}

	return data, nil
}

// TorrentTrackers method is used to get trackers of specified torrent.
// Pass hash of torrent as parameter so that server can identify the torrent.
func (client *Client) TorrentTrackers(hash string) ([]Tracker, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentTrackersEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrent trackers failed"})

	if err != nil {
		return nil, err
	}

	var data []Tracker
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
