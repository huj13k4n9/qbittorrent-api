package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	"strconv"
)

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
