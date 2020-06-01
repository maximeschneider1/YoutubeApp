package payload

import (
	"YoutubeApp/model"
	"fmt"
)

type User struct {
	ID        string
	Name      string
	URL       string
	Thumbnail string
}

//InfoFetcher regroups all methods to get info from Youtube API response.
// This interface regroups all behaviours of User object
type InfoFetcher interface {
	GetID(p model.Item) (string, error)
	GetURL(p model.Item) (string, error)
	GetTitle(p model.Item)(string, error)
	GetThumbnail(p model.Item) (string, error)
}

func (c *User) GetTitle(p model.Item) (string, error) {
	if p.Snippet.Title == "" {
		return "", fmt.Errorf("Couldn't get title")
	}
	title := p.Snippet.Title
	return title, nil
}
func (v *User) GetID(p model.Item) (string, error) {
	if p.Snippet.ResourceIDs.ChannelID == "" {
		return "", fmt.Errorf("Couldn't get title")
	}
	ID := p.Snippet.ResourceIDs.ChannelID
	return ID, nil
}
func (v *User) GetURL(p model.Item) (string, error){
	URL := fmt.Sprintf( `https://www.youtube.com/channel/%v`, p.Snippet.ResourceIDs.ChannelID)
	return URL, nil
}
func (c *User) GetThumbnail(p model.Item) (string, error) {
	thumbnail := p.Snippet.Thumbnails.Medium.URL
	return thumbnail, nil
}

//GetItemInfo gets infos from the Google response to
func (c *User)GetItemInfo(payload model.Item) (*User, error) {
	var err error
	c.ID, err = c.GetID(payload)
	if err != nil {
		return nil, err
	}
	c.URL, err = c.GetURL(payload)
	if err != nil {
		return nil, err
	}
	c.Name, err = c.GetTitle(payload)
	if err != nil {
		return nil, err
	}
	c.Thumbnail, err = c.GetThumbnail(payload)
	if err != nil {
		return nil, err
	}

	return c, nil
}