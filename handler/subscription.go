package handler

import (
	"YSS/model"
	"YSS/payload"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	AllSubscription []*payload.User
	AllVideosFromUser []*payload.Video
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	// Query result to the Youtube API
	var page Page
	subscriptions, err := querySubscription("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=10&mine=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	vidQuery := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?access_token=%v&channelId=UCC9mlCpyisiIpp9YA9xV-QA&part=snippet,id&maxResults=20", currentToken)

	response2, err := http.Get(vidQuery)
	defer response2.Body.Close()
	content2, err := ioutil.ReadAll(response2.Body)
	fmt.Printf("%#v",  string(content2))


	// Range over response items
	for _, p := range subscriptions.Items {
		c := &payload.User{}
		c, err = c.GetItemInfo(p); if err != nil {
			fmt.Println(err.Error())
		}
		page.AllSubscription = append(page.AllSubscription, c)
	}

	// Render results
	t, err := getTemplateHTML("./html/page.html"); if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, htmlHome)
	t.Execute(w, page.AllSubscription)
}


func HandleNextPage(w http.ResponseWriter, r *http.Request) {
	var page Page
	subscriptions, err := querySubscription("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=10&pageToken=CAoQAA&mine=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range subscriptions.Items {
		c := &payload.User{}

		c, err := c.GetItemInfo(p); if err != nil {
			fmt.Println(err.Error())
		}
		page.AllSubscription = append(page.AllSubscription, c)
	}

	t, err := getTemplateHTML("./html/page.html"); if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, htmlHome)
	t.Execute(w, page.AllSubscription)
}


func getTemplateHTML(filePath string) (*template.Template, error) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return &template.Template{}, err
	}
	return t, nil
}

func querySubscription(query string) (model.Payload, error) {
	queryBuild := fmt.Sprintf(query, currentToken)

	response, err := http.Get(queryBuild); if err != nil {
		return model.Payload{}, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	var result model.Payload
	err = json.Unmarshal(contents, &result); if err != nil {
		fmt.Println(err.Error())
	}
	return result, nil
}