package qbt

import "encoding/json"

type RSSRoot struct {
	Feeds    []*RSSData
	Folders  []*RSS
	Children []*RSS
}

type RSS struct {
	IsFolder bool
	Data     RSSData
	Children []*RSS
}

// RSSData Reference: https://github.com/qbittorrent/qBittorrent/blob/master/src/base/rss/rss_feed.cpp#L472
type RSSData struct {
	Name          string
	FullPath      string
	HasError      bool   `json:"hasError"`
	IsLoading     bool   `json:"isLoading"`
	LastBuildDate string `json:"lastBuildDate"`
	Title         string `json:"title"`
	UID           string `json:"uid"`
	URL           string `json:"url"`
	Articles      []*RSSArticle
}

// RSSArticle Reference: https://github.com/qbittorrent/qBittorrent/blob/master/src/base/rss/rss_parser.cpp#L604
type RSSArticle struct {
	ID          string         `json:"id"`
	Link        string         `json:"link"`
	Title       string         `json:"title"`
	TorrentURL  string         `json:"torrentURL"`
	Author      string         `json:"author"`
	Date        string         `json:"date"`
	Description string         `json:"description"`
	Other       map[string]any `json:"-"`
}

func (r *RSSArticle) UnmarshalJSON(bytes []byte) error {
	type Alias RSSArticle

	tmp := Alias{}

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	*r = RSSArticle(tmp)

	if err := json.Unmarshal(bytes, &r.Other); err != nil {
		return err
	}

	for _, v := range []string{"id", "link", "title", "torrentURL", "author", "date", "description"} {
		delete(r.Other, v)
	}

	return nil
}
