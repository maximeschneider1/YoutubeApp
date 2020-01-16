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
	AllSubscription []*payload.Channel
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	var page Page
	subscriptions, err := querySubscription("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=10&mine=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range subscriptions.Items {

		c := &payload.Channel{}

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
	querybuild := fmt.Sprintf(query, currentToken)

	response, err := http.Get(querybuild); if err != nil {
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

func HandlePageTwo(w http.ResponseWriter, r *http.Request) {
	var page Page
	subscriptions, err := querySubscription("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=10&pageToken=CAoQAA&mine=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range subscriptions.Items {

		c := &payload.Channel{}

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




