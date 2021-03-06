package model

import "time"

type Payload struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PrevPageToken string `json:"prevPageToken"`
	PageInfos      PageInfo `json:"pageInfo"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ResourceIDs ResourceID`json:"resourceId"`
	ChannelID  string `json:"channelId"`
	Thumbnails Thumbnail `json:"thumbnails"`
}

type ResourceID  struct {
Kind      string `json:"kind"`
ChannelID string `json:"channelId"`
}

type Thumbnail struct {
	Medium struct {
		Default struct {
			URL string `json:"url"`
		} `json:"default"`
		URL string `json:"url"`
	} `json:"medium"`
	High struct {
		URL string `json:"url"`
	} `json:"high"`
}