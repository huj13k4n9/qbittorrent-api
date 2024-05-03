package qbt

import (
	"encoding/json"
	"github.com/huj13k4n9/qbittorrent-api/consts"
	"strconv"
	"strings"
)

// BuildRSSTree is used by Client.GetAllRSSItems to build a tree
// structure of the RSS feeds and folders in qBittorrent.
func BuildRSSTree(input map[string]any, level int, path []string, root *RSSRoot, node *RSS) error {
	// Iterate each folder/feed
	for key, value := range input {
		// Record path of node
		path = append(path, key)

		if v, ok := value.(map[string]any); ok {
			// Check if map has `uid` and `url` attributes
			// If true, the value should be feed data
			// Otherwise, the value is recognized as a folder
			_, ok1 := v["uid"].(string)
			_, ok2 := v["url"].(string)
			if ok1 && ok2 {
				// Feed data
				feedData := RSSData{}
				err := MapToStruct(v, &feedData)
				if err != nil {
					return err
				}

				feedData.Name = key
				feedData.FullPath = strings.Join(path, "\\")
				feed := &RSS{
					IsFolder: false,
					Children: nil,
					Data:     feedData,
				}

				// Add reference in RSSRoot
				root.Feeds = append(root.Feeds, &feedData)

				// Build structure of the whole tree
				if level == 0 {
					root.Children = append(root.Children, feed)
				} else {
					node.Children = append(node.Children, feed)
				}
			} else {
				// Folder
				folder := &RSS{
					IsFolder: true,
					Data: RSSData{
						Name:     key,
						FullPath: strings.Join(path, "\\"),
					},
					Children: []*RSS{},
				}

				// Add reference in RSSRoot
				root.Folders = append(root.Folders, folder)

				// Build structure of the whole tree
				if level == 0 {
					root.Children = append(root.Children, folder)
				} else {
					node.Children = append(node.Children, folder)
				}

				// Recursive call, process the next layer
				err := BuildRSSTree(v, level+1, path, root, folder)
				if err != nil {
					return err
				}
			}
		}
		// Backtrace the path
		path = path[0 : len(path)-1]
	}
	return nil
}

// GetAllRSSItems This method is used to get all RSS items in
// qBittorrent, in a tree structure.
//
// Use RSSRoot.Feeds to quickly access all feeds and skip the
// structure of folders. Use RSSRoot.Folders to access all
// existing folders.
//
// Note that when RSS.IsFolder is true, only RSSData.Name and
// RSSData.FullPath are used, the others are ignored.
func (client *Client) GetAllRSSItems(withData bool) (*RSSRoot, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetAllRSSItemsEndpoint,
		map[string]string{"withData": strconv.FormatBool(withData)}, nil,
		map[string]string{"!200": "get RSS items failed"})

	if err != nil {
		return nil, err
	}

	var tempData map[string]any
	err = json.NewDecoder(resp.Body).Decode(&tempData)
	if err != nil {
		return nil, err
	}

	rootFolder := &RSSRoot{}
	err = BuildRSSTree(tempData, 0, []string{}, rootFolder, nil)
	if err != nil {
		return nil, err
	}

	return rootFolder, nil
}
