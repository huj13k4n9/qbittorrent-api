package qbt

type SearchStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type SearchResponse struct {
	Results []struct {
		DescriptionLink  string `json:"descrLink"`
		FileName         string `json:"fileName"`
		FileSize         int    `json:"fileSize"`
		FileURL          string `json:"fileUrl"`
		NumberOfLeechers int    `json:"nbLeechers"`
		NumberOfSeeders  int    `json:"nbSeeders"`
		SiteURL          string `json:"siteUrl"`
	} `json:"results"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type SearchPluginResult struct {
	Enabled             bool   `json:"enabled"`
	FullName            string `json:"fullName"`
	Name                string `json:"name"`
	SupportedCategories []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"supportedCategories"`
	Url     string `json:"url"`
	Version string `json:"version"`
}
