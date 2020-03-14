package handler

import (
	"YSS/model"
	"YSS/payload"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func random(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	var page Page

	// Build query
	randomOrderForQuery := randomOrder()

	baseQuery := "https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=50&order=%v&mine=true"
	queryBuild := fmt.Sprintf(baseQuery, currentToken, randomOrderForQuery)

	// Make query and unmarshall it
	response, err := http.Get(queryBuild); if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	var subscriptions model.Payload
	err = json.Unmarshal(contents, &subscriptions); if err != nil {
		fmt.Println(err.Error())
	}

	// Take the result
	// Range over response items
	for _, p := range subscriptions.Items {
		c := &payload.User{}
		c, err = c.GetItemInfo(p); if err != nil {
			fmt.Println(err.Error())
		}
		page.AllSubscription = append(page.AllSubscription, c)
	}

	winner := chooseItem2(page.AllSubscription)

	jsonBody, err := json.Marshal(winner)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}

	w.Write(jsonBody)
}

func chooseItem2(allUsers []*payload.User) *payload.User {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(allUsers)
	r := rand.Intn(max - min) + min

	var winner *payload.User

	winner = allUsers[r]

	return winner
}


// randomOrder returns random order param for the request
func randomOrder() string {
	rand.Seed(time.Now().UnixNano())
	//min := 1
	//max := 4
	r := rand.Intn(4)
	if r == 0 {
		randomOrder()
	}
	if r == 1 {
		return "alphabetical"
	}
	if r == 2 {
		return "relevance"
	}
	if r == 3 {
		return "unread"
	}
	return "alphabetical"
}