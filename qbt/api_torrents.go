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
// Pass hash of torrents as parameter so that server can identify them.
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
// Pass hash of torrents as parameter so that server can identify them.
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

// TorrentWebSeeds method is used to get web seeds of specified torrent.
// Pass hash of torrents as parameter so that server can identify them.
// Return URLs of web seeds as string.
func (client *Client) TorrentWebSeeds(hash string) ([]string, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentWebSeedsEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrent web seeds failed"})

	if err != nil {
		return nil, err
	}

	var data []struct {
		URL string `json:"url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, item := range data {
		ret = append(ret, item.URL)
	}

	return ret, nil
}

// TorrentContents method is used to get web seeds of specified torrent.
//
// Pass hash of torrents as parameter so that server can identify them.
// `indexes` is optional, which represents what indexes do you want to obtain
// file information with.
//
// Return a list of TorrentFileProperties, where each element contains info
// about one file.
func (client *Client) TorrentContents(hash string, indexes []int) ([]TorrentFileProperties, error) {
	params := make(map[string]string)
	params["hash"] = hash
	if indexes != nil && len(indexes) != 0 {
		params["indexes"] = func(array []int) string {
			var temp []string
			for _, item := range array {
				temp = append(temp, strconv.Itoa(item))
			}
			return strings.Join(temp, "|")
		}(indexes)
	}

	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentContentsEndpoint, params, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrent contents failed"})

	if err != nil {
		return nil, err
	}

	var data []TorrentFileProperties
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TorrentPieceStates method is used to get pieces' states of specified torrent.
//
// Pass hash of torrents as parameter so that server can identify them.
//
// Return an array of states (integers) of all pieces (in order) of a specific torrent
func (client *Client) TorrentPieceStates(hash string) ([]int, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentPieceStatesEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrent pieces' states failed"})

	if err != nil {
		return nil, err
	}

	var data []int
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TorrentPieceHashes method is used to get pieces' hashes of specified torrent.
//
// Pass hash of torrents as parameter so that server can identify them.
//
// Return an array of hashes (strings) of all pieces (in order) of a specific torrent
func (client *Client) TorrentPieceHashes(hash string) ([]string, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetTorrentPieceHashesEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{"404": "hash is invalid", "!200": "get torrent pieces' hashes failed"})

	if err != nil {
		return nil, err
	}

	var data []string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PauseTorrents method is used to pause specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) PauseTorrents(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.PauseTorrentsEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{"!200": "pause torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// ResumeTorrents method is used to resume specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) ResumeTorrents(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ResumeTorrentsEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{"!200": "resume torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// DeleteTorrents method is used to delete specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) DeleteTorrents(hashes []string, deleteFiles bool) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.DeleteTorrentsEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "deleteFiles": strconv.FormatBool(deleteFiles)},
		nil, map[string]string{"!200": "delete torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// RecheckTorrents method is used to recheck specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) RecheckTorrents(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RecheckTorrentsEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{"!200": "recheck torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// ReannounceTorrents method is used to reannounce specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) ReannounceTorrents(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ReannounceTorrentsEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{"!200": "reannounce torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// AddTrackersToTorrent method is used to add trackers to specified torrent.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) AddTrackersToTorrent(hash string, trackers []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.AddTrackersToTorrentEndpoint, map[string]string{"hash": hash, "trackers": strings.Join(trackers, "\n")},
		nil, map[string]string{"404": "torrent hash was not found", "!200": "add trackers to torrent failed"})

	if err != nil {
		return err
	}

	return nil
}

// RemoveTrackersToTorrent method is used to remove trackers to specified torrent.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) RemoveTrackersToTorrent(hash string, trackers []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RemoveTrackersEndpoint, map[string]string{"hash": hash, "trackers": strings.Join(trackers, "|")},
		nil, map[string]string{
			"404":  "torrent hash was not found",
			"409":  "specified trackers not found",
			"!200": "remove trackers to torrent failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// EditTrackersToTorrent method is used to edit trackers to specified torrent.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) EditTrackersToTorrent(hash string, origUrl string, newUrl string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.EditTrackersEndpoint, map[string]string{"hash": hash, "origUrl": origUrl, "newUrl": newUrl},
		nil, map[string]string{
			"400":  "newUrl is not a valid URL",
			"404":  "torrent hash was not found",
			"409":  "newUrl already exists for the torrent or origUrl was not found",
			"!200": "edit trackers to torrent failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// AddPeers method is used to add peers to specified torrent(s).
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) AddPeers(hashes []string, peers []Peer) error {
	var peerString []string
	for _, peer := range peers {
		peerString = append(peerString, peer.String())
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.AddPeersEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "peers": strings.Join(hashes, "|")},
		nil, map[string]string{
			"400":  "none of the supplied peers are valid",
			"!200": "add peers to torrent(s) failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// IncreaseTorrentPriority method is used to increase torrent priority.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) IncreaseTorrentPriority(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.IncreaseTorrentPriorityEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"409":  "torrent queueing is not enabled",
			"!200": "increase torrent priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// DecreaseTorrentPriority method is used to decrease torrent priority.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) DecreaseTorrentPriority(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.DecreaseTorrentPriorityEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"409":  "torrent queueing is not enabled",
			"!200": "decrease torrent priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// MaximalTorrentPriority method is used to maximize torrent priority.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) MaximalTorrentPriority(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.MaximalTorrentPriorityEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"409":  "torrent queueing is not enabled",
			"!200": "maximize torrent priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// MinimalTorrentPriority method is used to minimize torrent priority.
// Pass hash of torrents as parameter so that server can identify them.
func (client *Client) MinimalTorrentPriority(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.MinimalTorrentPriorityEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"409":  "torrent queueing is not enabled",
			"!200": "minimize torrent priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}
