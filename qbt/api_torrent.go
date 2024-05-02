package qbt

import (
	"bytes"
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// BuildTorrentListQuery is used to check request parameters
// of torrent list API, and then turn actual values into
// `map[string]string` for Client.Get to use.
func BuildTorrentListQuery(req *TorrentListParams) map[string]string {
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

// BuildAddTorrentsQuery is used to check request parameters
// of add new torrents API, and then turn actual values into
// a multipart form for Client.PostMultipart to use.
func BuildAddTorrentsQuery(req *AddTorrentParams, writer *multipart.Writer) error {
	// 1. params has to be valid
	// 2. one of params.TorrentFiles and params.TorrentURLs needs to be valid
	if !(req != nil &&
		((req.TorrentFiles != nil && len(req.TorrentFiles) != 0) ||
			(req.TorrentURLs != nil && len(req.TorrentURLs) != 0))) {
		return wrapper.Wrap(ErrUnknownType, "AddTorrentParams: mandatory parameters missing")
	}

	var err error
	if req.TorrentURLs != nil && len(req.TorrentURLs) != 0 {
		err = writer.WriteField("urls", strings.Join(req.TorrentURLs, "\r\n"))
		if err != nil {
			return err
		}
	}

	if req.TorrentFiles != nil && len(req.TorrentFiles) != 0 {
		for _, file := range req.TorrentFiles {
			fieldWriter, err := writer.CreateFormFile("torrents", filepath.Base(file))
			if err != nil {
				return err
			}

			fileReader, err := os.OpenFile(file, os.O_RDONLY, 0644)
			if err != nil {
				return err
			}

			_, err = io.Copy(fieldWriter, fileReader)
			if err != nil {
				return err
			}

			err = fileReader.Close()
			if err != nil {
				return err
			}
		}
	}

	if req.SavePath != "" {
		err = writer.WriteField("savepath", req.SavePath)
		if err != nil {
			return err
		}
	}

	if req.Cookie != "" {
		err = writer.WriteField("cookie", req.Cookie)
		if err != nil {
			return err
		}
	}

	if req.Category != "" {
		err = writer.WriteField("category", req.Category)
		if err != nil {
			return err
		}
	}

	if req.Tags != nil && len(req.Tags) != 0 {
		err = writer.WriteField("tags", strings.Join(req.Tags, ","))
		if err != nil {
			return err
		}
	}

	if req.Rename != "" {
		err = writer.WriteField("rename", req.Rename)
		if err != nil {
			return err
		}
	}

	if req.StopCondition != "" {
		err = writer.WriteField("contentLayout", req.StopCondition)
		if err != nil {
			return err
		}
	}

	if req.ContentLayout != "" {
		err = writer.WriteField("stopCondition", req.ContentLayout)
		if err != nil {
			return err
		}
	}

	if req.UploadLimit > 0 {
		err = writer.WriteField("upLimit", strconv.Itoa(req.UploadLimit))
		if err != nil {
			return err
		}
	}

	if req.DownloadLimit > 0 {
		err = writer.WriteField("dlLimit", strconv.Itoa(req.DownloadLimit))
		if err != nil {
			return err
		}
	}

	if req.RatioLimit > 0 {
		err = writer.WriteField("ratioLimit", strconv.FormatFloat(req.RatioLimit, 'g', -1, 64))
		if err != nil {
			return err
		}
	}

	if req.SeedingTimeLimit > 0 {
		err = writer.WriteField("seedingTimeLimit", strconv.Itoa(req.SeedingTimeLimit))
		if err != nil {
			return err
		}
	}

	if req.SeedingTimeLimit > 0 {
		err = writer.WriteField("seedingTimeLimit", strconv.Itoa(req.SeedingTimeLimit))
		if err != nil {
			return err
		}
	}

	err = writer.WriteField("skip_checking", strconv.FormatBool(req.SkipChecking))
	if err != nil {
		return err
	}

	err = writer.WriteField("paused", strconv.FormatBool(req.Paused))
	if err != nil {
		return err
	}

	err = writer.WriteField("autoTMM", strconv.FormatBool(req.AutoTMM))
	if err != nil {
		return err
	}

	err = writer.WriteField("sequentialDownload", strconv.FormatBool(req.SequentialDownload))
	if err != nil {
		return err
	}

	err = writer.WriteField("firstLastPiecePrio", strconv.FormatBool(req.FirstLastPiecePrioritized))
	if err != nil {
		return err
	}

	err = writer.WriteField("addToTopOfQueue", strconv.FormatBool(req.AddToTopOfQueue))
	if err != nil {
		return err
	}

	switch req.CreateRootFolder {
	case "true":
	case "false":
		err = writer.WriteField("root_folder", req.CreateRootFolder)
		if err != nil {
			return err
		}
		break
	case "unset":
	default:
		break
	}

	return nil
}

// Torrents method is used to get torrent list in qBittorrent.
// Return basic info of all torrents in TorrentInfo struct.
//
// If you want to request directly with no parameters, pass
// options with nil. Otherwise, construct your own TorrentListParams
// data is needed.
func (client *Client) Torrents(options *TorrentListParams) ([]TorrentInfo, error) {
	var params map[string]string
	if options == nil {
		params = BuildTorrentListQuery(&TorrentListParams{
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
func (client *Client) ReannounceTorrents(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ReannounceTorrentsEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{"!200": "reannounce torrents failed"})

	if err != nil {
		return err
	}

	return nil
}

// ExportTorrent method is used to export an existing torrent.
func (client *Client) ExportTorrent(hash string) ([]byte, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.ExportTorrentEndpoint, map[string]string{"hash": hash}, nil,
		map[string]string{
			"404":  "torrent hash was not found",
			"!200": "export torrent failed",
		})

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExportTorrentToFile method is used to export an existing torrent
// to a specified file.
//
// `overwrite` parameter controls whether overwrite original
// file when file exists. If it's set to `false`, os.O_EXCL will be
// used. Otherwise, os.O_TRUNC will be used.
func (client *Client) ExportTorrentToFile(hash string, location string, overwrite bool) error {
	torrent, err := client.ExportTorrent(hash)
	if err != nil {
		return err
	}

	err = WriteFile(location, torrent, overwrite)
	if err != nil {
		return err
	}

	return nil
}

// AddNewTorrents method is used to add new torrent to qBittorrent.
func (client *Client) AddNewTorrents(params *AddTorrentParams) error {
	if !client.Authenticated {
		return ErrUnauthenticated
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	err := BuildAddTorrentsQuery(params, writer)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	resp, err := client.PostMultipart(consts.AddNewTorrentEndpoint, &requestBody, writer.FormDataContentType())
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnsupportedMediaType:
		return wrapper.Wrap(ErrBadResponse, "torrent file is not valid")
	default:
		return wrapper.Wrap(ErrBadResponse, "add new torrent failed")
	}
}

// AddTrackersToTorrent method is used to add trackers to specified torrent.
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
func (client *Client) AddPeers(hashes []string, peers []Peer) error {
	var peerString []string
	for _, peer := range peers {
		peerString = append(peerString, peer.String())
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.AddPeersEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "peers": strings.Join(peerString, "|")},
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

// SetFilePriority method is used to set files' priority inside a torrent.
// Use indexes from TorrentFileProperties.Index to identify files.
func (client *Client) SetFilePriority(hash string, indexes []int, priority int) error {
	var indexString []string
	for _, index := range indexes {
		indexString = append(indexString, strconv.Itoa(index))
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.SetFilePriorityEndpoint,
		map[string]string{
			"hash":     hash,
			"indexes":  strings.Join(indexString, "|"),
			"priority": strconv.Itoa(priority),
		}, nil,
		map[string]string{
			"400":  "priority is invalid, or at least one file id is not a valid integer",
			"404":  "torrent hash was not found",
			"409":  "torrent metadata hasn't downloaded yet, or at least one file id was not found",
			"!200": "set file priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// GetDownloadLimit method is used to get download speed limit of torrent(s).
func (client *Client) GetDownloadLimit(hashes []string) (map[string]int, error) {
	resp, err := client.RequestAndHandleError(
		"POST", consts.GetTorrentDownloadLimitEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"!200": "get download limit failed",
		})

	if err != nil {
		return nil, err
	}

	var data map[string]int
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUploadLimit method is used to get upload speed limit of torrent(s).
func (client *Client) GetUploadLimit(hashes []string) (map[string]int, error) {
	resp, err := client.RequestAndHandleError(
		"POST", consts.GetTorrentUploadLimitEndpoint, map[string]string{"hashes": strings.Join(hashes, "|")},
		nil, map[string]string{
			"!200": "get upload limit failed",
		})

	if err != nil {
		return nil, err
	}

	var data map[string]int
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SetDownloadLimit method is used to set download speed limit of torrent(s).
func (client *Client) SetDownloadLimit(hashes []string, limit int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentDownloadLimitEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "limit": strconv.Itoa(limit)}, nil,
		map[string]string{
			"!200": "set download limit failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetUploadLimit method is used to set upload speed limit of torrent(s).
func (client *Client) SetUploadLimit(hashes []string, limit int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentUploadLimitEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "limit": strconv.Itoa(limit)}, nil,
		map[string]string{
			"!200": "set upload limit failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetShareLimit method is used to set share limit of torrent(s),
// including ratio limit and seeding time limit.
//
// Hashes can contain multiple hashes separated by `|` or set to `all`.
// ratioLimit is the max ratio the torrent should be seeded until.
// `-2` means the global limit should be used, `-1` means no limit.
// seedingTimeLimit is the max amount of time (minutes) the torrent
// should be seeded. `-2` means the global limit should be used,
// `-1` means no limit.
func (client *Client) SetShareLimit(hashes []string, ratioLimit float64, seedingTimeLimit int) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentShareLimitEndpoint,
		map[string]string{
			"hashes":           strings.Join(hashes, "|"),
			"ratioLimit":       strconv.FormatFloat(ratioLimit, 'g', -1, 64),
			"seedingTimeLimit": strconv.Itoa(seedingTimeLimit),
		}, nil,
		map[string]string{
			"!200": "set share limit failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetTorrentLocation method is used to location to download the torrent to.
func (client *Client) SetTorrentLocation(hashes []string, location string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentLocationEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "location": location}, nil,
		map[string]string{
			"400":  "save path is empty",
			"403":  "user does not have write access to directory",
			"409":  "unable to create save path directory",
			"!200": "set torrent location failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetTorrentName method is used to set name of torrent.
func (client *Client) SetTorrentName(hash string, name string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentNameEndpoint,
		map[string]string{"hash": hash, "name": name}, nil,
		map[string]string{
			"404":  "torrent hash is invalid",
			"409":  "torrent name is empty",
			"!200": "set torrent name failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetTorrentCategory method is used to set category of torrent.
func (client *Client) SetTorrentCategory(hashes []string, category string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetTorrentCategoryEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "category": category}, nil,
		map[string]string{
			"409":  "category name does not exist",
			"!200": "set torrent category failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// AddTorrentTags method is used to add tags of torrent.
func (client *Client) AddTorrentTags(hashes []string, tags []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.AddTorrentTagsEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "tags": strings.Join(tags, ",")}, nil,
		map[string]string{
			"!200": "add torrent tags failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// RemoveTorrentTags method is used to remove tags of torrent.
func (client *Client) RemoveTorrentTags(hashes []string, tags []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RemoveTorrentTagsEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "tags": strings.Join(tags, ",")}, nil,
		map[string]string{
			"!200": "remove torrent category failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// Categories method is used to get all categories in qBittorrent.
func (client *Client) Categories() ([]Category, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetAllCategoriesEndpoint, nil, nil,
		map[string]string{
			"!200": "get categories failed",
		})

	if err != nil {
		return nil, err
	}

	var temp map[string]Category
	var data []Category
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return nil, err
	}

	for _, v := range temp {
		data = append(data, v)
	}

	return data, nil
}

// AddCategory method is used to add a category to qBittorrent.
func (client *Client) AddCategory(category Category) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.AddNewCategoryEndpoint, map[string]string{
			"category": category.Name,
			"savePath": category.SavePath,
		}, nil,
		map[string]string{
			"400":  "category name is empty",
			"409":  "category name is invalid",
			"!200": "add category failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// EditCategory method is used to edit a category to qBittorrent.
func (client *Client) EditCategory(category Category) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.EditCategoryEndpoint, map[string]string{
			"category": category.Name,
			"savePath": category.SavePath,
		}, nil,
		map[string]string{
			"400":  "category name is empty",
			"409":  "category editing failed",
			"!200": "category editing failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// RemoveCategories method is used to remove categories in qBittorrent.
func (client *Client) RemoveCategories(categories []Category) error {
	categoryString := make([]string, len(categories))
	for _, category := range categories {
		categoryString = append(categoryString, category.Name)
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.EditCategoryEndpoint, map[string]string{
			"categories": strings.Join(categoryString, "\n"),
		}, nil,
		map[string]string{
			"!200": "remove categories failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// RemoveCategoriesWithNames method is used to remove categories in qBittorrent.
func (client *Client) RemoveCategoriesWithNames(categories []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RemoveCategoriesEndpoint, map[string]string{
			"categories": strings.Join(categories, "\n"),
		}, nil,
		map[string]string{
			"!200": "remove categories failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// Tags method is used to get all tags in qBittorrent.
func (client *Client) Tags() ([]string, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetAllTagsEndpoint, nil, nil,
		map[string]string{
			"!200": "get tags failed",
		})

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

// CreateTags method is used to create new tags in qBittorrent.
func (client *Client) CreateTags(tags []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.CreateTagsEndpoint,
		map[string]string{"tags": strings.Join(tags, ",")}, nil,
		map[string]string{
			"!200": "create tags failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// DeleteTags method is used to delete existing tags in qBittorrent.
func (client *Client) DeleteTags(tags []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.DeleteTagsEndpoint,
		map[string]string{"tags": strings.Join(tags, ",")}, nil,
		map[string]string{
			"!200": "delete tags failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetAutoTorrentManagement method is used to set auto torrent management mode of torrent(s).
func (client *Client) SetAutoTorrentManagement(hashes []string, enable bool) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetAutoTorrentManagementEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "enable": strconv.FormatBool(enable)}, nil,
		map[string]string{
			"!200": "set auto torrent management failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// ToggleSequentialDownload method is used to toggle sequential download mode of torrent(s).
func (client *Client) ToggleSequentialDownload(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.ToggleSequentialDownloadEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|")}, nil,
		map[string]string{
			"!200": "toggle sequential download failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetFirstLastPiecePriority method is used to toggle the first/last piece priority of torrent(s).
func (client *Client) SetFirstLastPiecePriority(hashes []string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetFirstLastPiecePriorityEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|")}, nil,
		map[string]string{
			"!200": "set first/last piece priority failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetForceStart method is used to set force start mode of given torrent(s).
func (client *Client) SetForceStart(hashes []string, enable bool) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetForceStartEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "value": strconv.FormatBool(enable)}, nil,
		map[string]string{
			"!200": "set force start failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// SetSuperSeeding method is used to set super seeding mode of given torrent(s).
func (client *Client) SetSuperSeeding(hashes []string, enable bool) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.SetSuperSeedingEndpoint,
		map[string]string{"hashes": strings.Join(hashes, "|"), "value": strconv.FormatBool(enable)}, nil,
		map[string]string{
			"!200": "set super seeding failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// RenameFile method is used to rename file inside given torrent.
func (client *Client) RenameFile(hash string, oldPath string, newPath string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RenameFileEndpoint,
		map[string]string{"hash": hash, "oldPath": oldPath, "newPath": newPath}, nil,
		map[string]string{
			"400":  "missing newPath parameter",
			"409":  "invalid newPath or oldPath, or newPath is already in use",
			"!200": "rename file failed",
		})

	if err != nil {
		return err
	}

	return nil
}

// RenameFolder method is used to rename folder inside given torrent.
func (client *Client) RenameFolder(hash string, oldPath string, newPath string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RenameFolderEndpoint,
		map[string]string{"hash": hash, "oldPath": oldPath, "newPath": newPath}, nil,
		map[string]string{
			"400":  "missing newPath parameter",
			"409":  "invalid newPath or oldPath, or newPath is already in use",
			"!200": "rename folder failed",
		})

	if err != nil {
		return err
	}

	return nil
}
