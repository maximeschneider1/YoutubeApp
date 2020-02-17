package handler

import (
	"YSS/model"
	"YSS/payload"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleRandom(w http.ResponseWriter, r *http.Request) {
	// Query result to the Youtube API
	var page Page
	subscriptions, err := querySubscription("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=3&mine=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	var allIDs []string
	// Range over response items
	for _, p := range subscriptions.Items {
		c := &payload.User{}
		c, err = c.GetItemInfo(p); if err != nil {
			fmt.Println(err.Error())
		}
		page.AllSubscription = append(page.AllSubscription, c)
		allIDs = append(allIDs, c.ID)
	}

	// For every channels IDs search for X video
	for _, v := range allIDs {
		vidQuery := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?access_token=%v&channelId=%v&part=snippet,id&maxResults=1", currentToken, v)
		response2, err := http.Get(vidQuery); if err != nil {
			fmt.Println(err.Error())
		}
		defer response2.Body.Close()
		content2, _ := ioutil.ReadAll(response2.Body)
		var videoSearch model.VideoSearch
		err = json.Unmarshal(content2, &videoSearch); if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Fprintf(w, string(content2))

		// For every videos, append payload to page.AllVideosFromUser object
		for _, x := range videoSearch.Items {
			n := &payload.Video{
				ID:        x.IDs.VideoID,
				Title:     x.Snippet.Title,
				URL:       fmt.Sprintf( `https://www.youtube.com/watch?v=%v`, x.IDs.VideoID),
				Thumbnail: x.Snippet.Thumbnails.High.URL,
				Author: x.Snippet.ChannelTitle,
			}
			page.AllVideosFromUser = append(page.AllVideosFromUser, n)
		}
	}

	// Render results
	t, err := getTemplateHTML("./html/videoSearch.html"); if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, htmlHome)
	t.Execute(w, page.AllVideosFromUser)
}
