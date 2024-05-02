package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	"strconv"
)

// SyncMain The main sync function API. Used to sync the
// real-time status of qBittorrent. RID is used as a sequence
// of responses, and every response with bigger RID is a
// delta base on previous responses.
func (client *Client) SyncMain(rid int) (SyncMainData, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.SyncMainDataEndpoint, map[string]string{"rid": strconv.Itoa(rid)}, nil,
		map[string]string{"!200": "get sync main data failed"},
	)

	if err != nil {
		return SyncMainData{}, err
	}

	var data SyncMainData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return SyncMainData{}, err
	}

	return data, nil
}

// SyncTorrentPeers is a function used to sync real-time
// status of torrent peers in qBittorrent. RID is used
// as a sequence of responses, and every response with
// bigger RID is a delta base on previous responses.
func (client *Client) SyncTorrentPeers(hash string, rid int) (SyncPeersData, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.TorrentPeersDataEndpoint, map[string]string{"rid": strconv.Itoa(rid), "hash": hash}, nil,
		map[string]string{"404": "torrent hash was not found", "!200": "get sync peers data failed"},
	)

	if err != nil {
		return SyncPeersData{}, err
	}

	var data SyncPeersData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return SyncPeersData{}, err
	}

	return data, nil
}
