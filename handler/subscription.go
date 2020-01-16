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
	subscriptions, err := querySubscription()
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
	t2, err := getTemplateHTML("./html/page2.html"); if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprintf(w, htmlHome)
	t.Execute(w, page.AllSubscription[0])
	t2.Execute(w, page.AllSubscription[1])

}


func getTemplateHTML(filePath string) (*template.Template, error) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return &template.Template{}, err
	}
	return t, nil
}

func querySubscription() (model.Payload, error) {
	querybuild := fmt.Sprintf("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=2&mine=true", currentToken)
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
	querybuild2 := fmt.Sprintf("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=2&pageToken=CAIQAA&mine=true", currentToken)
	response2, err := http.Get(querybuild2); if err != nil {
		return
	}
	defer response2.Body.Close()
	contents2, err := ioutil.ReadAll(response2.Body)
	response2.Body.Close()
	fmt.Fprintf(w, htmlHome)
	fmt.Fprintf(w, "\n \n Content 2: %s\n", contents2)
	fmt.Fprintf(w, htmlHome)
}




