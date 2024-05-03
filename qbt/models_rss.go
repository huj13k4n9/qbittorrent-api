package qbt

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

type RSSArticle struct {
	Author      string `json:"author"`
	Creator     string `json:"creator"`
	Category    string `json:"category"`
	Date        string `json:"date"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	TorrentURL  string `json:"torrentURL"`
	CommentRSS  string `json:"commentRss"`
	Comments    string `json:"comments"`
}
