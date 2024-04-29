package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"net/http"
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
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

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

	resp, err := client.Get(consts.GetTorrentListEndpoint, params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, wrapper.Wrap(ErrBadResponse, "get torrents list failed")
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
	if !client.Authenticated {
		return TorrentProperties{}, ErrUnauthenticated
	}

	resp, err := client.Get(consts.GetTorrentPropertiesEndpoint, map[string]string{"hash": hash})
	if err != nil {
		return TorrentProperties{}, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		return TorrentProperties{}, wrapper.Wrap(ErrBadResponse, "hash is invalid")
	default:
		return TorrentProperties{}, wrapper.Wrap(ErrBadResponse, "get torrent properties failed")
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
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	resp, err := client.Get(consts.GetTorrentTrackersEndpoint, map[string]string{"hash": hash})
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		return nil, wrapper.Wrap(ErrBadResponse, "hash is invalid")
	default:
		return nil, wrapper.Wrap(ErrBadResponse, "get torrent trackers failed")
	}

	var data []Tracker
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
