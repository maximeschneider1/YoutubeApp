package payload

import (
	"YoutubeApp/model"
	"fmt"
)

type Video struct {
	ID    string
	Title string
	URL   string
	Thumbnail string
	Author string 
}

type VideoInfoFetcher interface {
	//GetID(p model.Item) (string, error)
	GetURL(s model.SearchItem) (string, error)
	GetTitle(s model.SearchItem)(string, error)
	//GetThumbnail(p model.Item) (string, error)
}

func (v *Video) GetTitle(s model.SearchItem) (string, error) {
	if s.Snippet.Title == "" {
		return "", fmt.Errorf("Couldn't get title")
	}
	title := s.Snippet.Title
	return title, nil
}

func (v *Video) GetURL(s model.SearchItem) (string, error) {
	if s.IDs.VideoID == "" {
		return "", fmt.Errorf("Couldn't get video Id")
	}
	ID := s.IDs.VideoID
	return ID, nil
}