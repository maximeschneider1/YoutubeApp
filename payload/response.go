package payload

import (
	"YSS/model"
	"fmt"
)

type Channel struct {
	ID    string
	Title string
	URL   string
	Thumbnail string
}

//InfoFetcher regroups all methods to get video info from Youtube API response.
// This interface regroups all behaviours of Channel object
type InfoFetcher interface {
	GetItemInfo(payload model.Item) (*Channel, error)

	GetID(p model.Item) (string, error)
	GetURL(p model.Item) (string, error)
	GetTitle(p model.Item)(string, error)
	GetThumbnail(p model.Item) (string, error)
}

//GetItemInfo gets Channel infos from the Google response
func (c *Channel)GetItemInfo(payload model.Item) (*Channel, error) {
	var err error
	c.ID, err = c.GetID(payload)
	if err != nil {
		return nil, err
	}
	c.URL, err = c.GetURL(payload)
	if err != nil {
		return nil, err
	}
	c.Title, err = c.GetTitle(payload)
	if err != nil {
		return nil, err
	}
	c.Thumbnail, err = c.GetThumbnail(payload)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Channel) GetTitle(p model.Item) (string, error) {
	if p.Snippet.Title == "" {
		return "", fmt.Errorf("Couldn't get title")
	}
	title := p.Snippet.Title
	return title, nil
}
func (v *Channel) GetID(p model.Item) (string, error) {
	if p.Snippet.ResourceIDs.ChannelID == "" {
		return "", fmt.Errorf("Couldn't get title")
	}
	ID := p.Snippet.ResourceIDs.ChannelID
	return ID, nil
}
func (v *Channel) GetURL(p model.Item) (string, error){
	URL := fmt.Sprintf( `https://www.youtube.com/channel/%v`, p.Snippet.ResourceIDs.ChannelID)
	return URL, nil
}
func (c *Channel) GetThumbnail(p model.Item) (string, error) {
	thumbnail := p.Snippet.Thumbnails.Medium.URL
	return thumbnail, nil
}