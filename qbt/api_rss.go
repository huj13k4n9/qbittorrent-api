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

// MoveRSSItem method is used to move a RSS item (feed or
// folder) in qBittorrent RSS module.
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) MoveRSSItem(src string, dst string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.MoveRSSItemEndpoint,
		map[string]string{"itemPath": src, "destPath": dst}, nil,
		map[string]string{"!200": "move RSS item failed"})

	if err != nil {
		return err
	}

	return nil
}

// RemoveRSSItem method is used to remove a RSS item (feed or
// folder) in qBittorrent RSS module.
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) RemoveRSSItem(path string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RemoveRSSItemEndpoint,
		map[string]string{"path": path}, nil,
		map[string]string{"!200": "remove RSS item failed"})

	if err != nil {
		return err
	}

	return nil
}

// AddRSSFeed method can add a new feed in qBittorrent RSS module.
// Parameter `path` is optional, when not provided, the feed will
// be stored in root folder.
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) AddRSSFeed(feed string, path string) error {
	params := map[string]string{"url": feed}

	if path != "" {
		params["path"] = path
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.AddRSSFeedEndpoint,
		params, nil,
		map[string]string{"!200": "add RSS feed failed"})

	if err != nil {
		return err
	}

	return nil
}

// AddRSSFolder method can add a folder in qBittorrent RSS module.
// RSS feeds can be stored in a specific folder.
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) AddRSSFolder(path string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.AddRSSFolderEndpoint,
		map[string]string{"path": path}, nil,
		map[string]string{"!200": "add RSS folder failed"})

	if err != nil {
		return err
	}

	return nil
}

// MarkAsRead method is used to mark article status as read.
//
// If `articleId` is provided only the article is marked as read
// otherwise the whole feed is going to be marked as read
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) MarkAsRead(path string, articleId string) error {
	params := map[string]string{"itemPath": path}

	if articleId != "" {
		params["articleId"] = articleId
	}

	_, err := client.RequestAndHandleError(
		"POST", consts.MarkAsReadEndpoint, params, nil,
		map[string]string{"!200": "mark as read failed"})

	if err != nil {
		return err
	}

	return nil
}

// RefreshRSSItem method is used to refresh the content of a
// RSS item (feed or folder) in qBittorrent RSS module.
//
// Path of item should use `\` as delimiter instead of `/` or
// anything else.
func (client *Client) RefreshRSSItem(path string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RefreshRSSItemEndpoint,
		map[string]string{"path": path}, nil,
		map[string]string{"!200": "refresh RSS item failed"})

	if err != nil {
		return err
	}

	return nil
}

// GetAllAutoDownloadRules method is used to get all auto-downloading
// rules of RSS module in qBittorrent. For definition of an
// auto-downloading rule, refer to type AutoDownloadRule.
func (client *Client) GetAllAutoDownloadRules() ([]*AutoDownloadRule, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.GetAllAutoDownloadRulesEndpoint,
		nil, nil,
		map[string]string{"!200": "get auto download rules failed"})

	if err != nil {
		return nil, err
	}

	var temp map[string]AutoDownloadRule
	var data []*AutoDownloadRule
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return nil, err
	}

	for k, v := range temp {
		v.Name = k
		data = append(data, &v)
	}

	return data, nil
}

// SetAutoDownloadRule method is used to add a new auto-downloading
// rule of RSS module in qBittorrent. For definition of an
// auto-downloading rule, refer to type AutoDownloadRule.
func (client *Client) SetAutoDownloadRule(rule *AutoDownloadRule) error {
	bytes, err := json.Marshal(*rule)
	if err != nil {
		return err
	}

	_, err = client.RequestAndHandleError(
		"POST", consts.SetAutoDownloadRuleEndpoint,
		map[string]string{"ruleName": rule.Name, "ruleDef": string(bytes)}, nil,
		map[string]string{"!200": "set auto download rule failed"})

	if err != nil {
		return err
	}

	return nil
}

// RenameAutoDownloadRule method is used to rename an existing
// auto-downloading rule of RSS module in qBittorrent. For
// definition of an auto-downloading rule, refer to type
// AutoDownloadRule.
func (client *Client) RenameAutoDownloadRule(ruleName string, newRuleName string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RenameAutoDownloadRuleEndpoint,
		map[string]string{"ruleName": ruleName, "newRuleName": newRuleName}, nil,
		map[string]string{"!200": "rename auto download rule failed"})

	if err != nil {
		return err
	}

	return nil
}

// RemoveAutoDownloadRule method is used to remove an existing
// auto-downloading rule of RSS module in qBittorrent. For
// definition of an auto-downloading rule, refer to type
// AutoDownloadRule.
func (client *Client) RemoveAutoDownloadRule(ruleName string) error {
	_, err := client.RequestAndHandleError(
		"POST", consts.RemoveAutoDownloadRuleEndpoint,
		map[string]string{"ruleName": ruleName}, nil,
		map[string]string{"!200": "remove auto download rule failed"})

	if err != nil {
		return err
	}

	return nil
}

// GetRuleMatchingArticles method is used to get all RSS
// articles matched by a specific rule. Return all matched
// names of articles, associated with their feed name.
func (client *Client) GetRuleMatchingArticles(ruleName string) ([]*RuleMatchResult, error) {
	resp, err := client.RequestAndHandleError(
		"GET", consts.MatchArticlesWithRuleEndpoint,
		map[string]string{"ruleName": ruleName}, nil,
		map[string]string{"!200": "failed to get all articles matching a rule"})

	if err != nil {
		return nil, err
	}

	var tmp map[string][]string
	var data []*RuleMatchResult
	err = json.NewDecoder(resp.Body).Decode(&tmp)
	if err != nil {
		return nil, err
	}

	for k, v := range tmp {
		data = append(data, &RuleMatchResult{
			FeedName:     k,
			ArticleNames: v,
		})
	}

	return data, nil
}
