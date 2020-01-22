package model

import "time"

type VideoSearch struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfos      PageInfo `json:"pageInfo"`
	Items []SearchItem `json:"items"`
}
type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type SearchItem struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	IDs   ID `json:"id"`
	Snippet SearchSnippet `json:"snippet"`
}

type ID   struct {
Kind    string `json:"kind"`
VideoID string `json:"videoId"`
}

type SearchSnippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelID   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnails  SearchThumbnails `json:"thumbnails"`
	ChannelTitle         string `json:"channelTitle"`
	LiveBroadcastContent string `json:"liveBroadcastContent"`
}

type SearchThumbnails  struct {
	Default struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"default"`
	Medium struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"medium"`
	High struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"high"`
}
