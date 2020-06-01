package handler

import (
	"YoutubeApp/model"
	"YoutubeApp/payload"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)



type Page struct {
	AllSubscription []*payload.User
	AllVideosFromUser []*payload.Video
}

type Claims struct {
	Username string `json:"userID"`
	OauthToken string `json:"oauthToken"`
	jwt.StandardClaims
}

func (s *server)HandleGetSubs() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")
		if r.Method != "GET" {
			return
		}

		tokenFromHeader := r.Header.Get("jwtToken")
		claims := &Claims{}
		_, err := jwt.ParseWithClaims(tokenFromHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		}); if err != nil {
			fmt.Println("Error parsing JWT token claims:", err.Error())
			jsonBody, err := json.Marshal(err.Error())
			if err != nil {
				http.Error(w, "Error converting results to json",
					http.StatusInternalServerError)
			}
			w.WriteHeader(403)
			w.Write(jsonBody)
			return
		}
		userOauthToken := claims.OauthToken

		nextPageToken := r.Header.Get("nextPageToken")
		prevPageToken := r.Header.Get("prevPageToken")

		if nextPageToken != "" {
			jsonBody := nextSubs(w, userOauthToken, nextPageToken)
			w.Write(jsonBody)
			return
		}
		if prevPageToken != "" {
			jsonBody := nextSubs(w, userOauthToken, prevPageToken)
			w.Write(jsonBody)
			return
		}

		// last case
		if prevPageToken != "" && nextPageToken == "" {

			// Query result to the Youtube API
			var page Page
			subscriptions, err := queryNextSubs("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=12&pageToken=%v&mine=true", userOauthToken, prevPageToken)
			if err != nil {
				fmt.Println("error querying user subscriptions :", err.Error())
			}
			// Range over response items
			for _, p := range subscriptions.Items {
				c := &payload.User{}
				c, err = c.GetItemInfo(p); if err != nil {
					fmt.Println("Error retrieving items information :",err.Error())
				}
				page.AllSubscription = append(page.AllSubscription, c)
			}

			w.Header().Set("nextPageToken", subscriptions.NextPageToken)

			type resp struct {
				Subscriptions []*payload.User
				NextPageToken string
				PrevPageToken string
				TotalResults int
				ResultPerPage int
			}

			response := &resp{
				Subscriptions: page.AllSubscription,
				NextPageToken: subscriptions.NextPageToken,
				PrevPageToken: subscriptions.PrevPageToken,
				TotalResults: subscriptions.PageInfos.TotalResults,
				ResultPerPage: subscriptions.PageInfos.ResultsPerPage,
			}

			jsonBody, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error converting results to json",
					http.StatusInternalServerError)
			}

			w.Write(jsonBody)

			return
		}

		if nextPageToken != "" && prevPageToken != "" {

			// Query result to the Youtube API
			var page Page
			subscriptions, err := queryNextSubs("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=12&pageToken=%v&mine=true", userOauthToken, nextPageToken)
			if err != nil {
				fmt.Println("error querying user subscriptions :", err.Error())
			}
			// Range over response items
			for _, p := range subscriptions.Items {
				c := &payload.User{}
				c, err = c.GetItemInfo(p); if err != nil {
					fmt.Println("Error retrieving items information :",err.Error())
				}
				page.AllSubscription = append(page.AllSubscription, c)
			}

			w.Header().Set("nextPageToken", subscriptions.NextPageToken)

			type resp struct {
				Subscriptions []*payload.User
				NextPageToken string
				PrevPageToken string
				TotalResults int
				ResultPerPage int
			}

			response := &resp{
				Subscriptions: page.AllSubscription,
				NextPageToken: subscriptions.NextPageToken,
				PrevPageToken: subscriptions.PrevPageToken,
				TotalResults: subscriptions.PageInfos.TotalResults,
				ResultPerPage: subscriptions.PageInfos.ResultsPerPage,
			}

			jsonBody, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error converting results to json",
					http.StatusInternalServerError)
			}

			w.Write(jsonBody)

			return
		}

		// Query result to the Youtube API
		var page Page
		subscriptions, err := querySubs("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=12&mine=true", userOauthToken)
		if err != nil {
			fmt.Println("error querying user subscriptions :", err.Error())
		}

		// Range over response items
		for _, p := range subscriptions.Items {
			c := &payload.User{}
			c, err = c.GetItemInfo(p); if err != nil {
				fmt.Println("Error retrieving items information :",err.Error())
			}
			page.AllSubscription = append(page.AllSubscription, c)
		}

		w.Header().Set("nextPageToken", subscriptions.NextPageToken)

		type resp struct {
			Subscriptions []*payload.User
			NextPageToken string
			PrevPageToken string
			TotalResults int
			ResultPerPage int
		}

		response := &resp{
			Subscriptions: page.AllSubscription,
			NextPageToken: subscriptions.NextPageToken,
			TotalResults: subscriptions.PageInfos.TotalResults,
			ResultPerPage: subscriptions.PageInfos.ResultsPerPage,
			PrevPageToken: "",
		}

		jsonBody, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}

		w.Write(jsonBody)
	}
}

func (s *server)HandleNextSubs() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, jwtToken, nextPageToken, prevPageToken")

		if r.Method != "POST" {
			return
		}

		decoder := json.NewDecoder(r.Body)

		type myData struct {
			Salut string
		}

		var data myData
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func queryNextSubs(query string, oauthToken string, nextPageToken string) (model.Payload, error) {
	queryBuild := fmt.Sprintf(query, oauthToken, nextPageToken)

	response, err := http.Get(queryBuild); if err != nil {
		return model.Payload{}, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	var result model.Payload
	err = json.Unmarshal(contents, &result); if err != nil {
		return model.Payload{}, err
	}
	return result, nil
}

func querySubs(query string, oauthToken string) (model.Payload, error) {
	queryBuild := fmt.Sprintf(query, oauthToken)

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

func nextSubs(w http.ResponseWriter, pageToken string, userOauthToken string) []byte {
	// Query result to the Youtube API
	var page Page
	subscriptions, err := queryNextSubs("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=16&pageToken=%v&mine=true", pageToken, userOauthToken)
	if err != nil {
		fmt.Println("error querying user subscriptions :", err.Error())
	}
	// Range over response items
	for _, p := range subscriptions.Items {
		c := &payload.User{}
		c, err = c.GetItemInfo(p); if err != nil {
			fmt.Println("Error retrieving items information :",err.Error())
		}
		page.AllSubscription = append(page.AllSubscription, c)
	}

	w.Header().Set("nextPageToken", subscriptions.NextPageToken)

	type resp struct {
		Subscriptions []*payload.User
		NextPageToken string
		PrevPageToken string
		TotalResults int
		ResultPerPage int
	}

	response := &resp{
		Subscriptions: page.AllSubscription,
		NextPageToken: subscriptions.NextPageToken,
		PrevPageToken: subscriptions.PrevPageToken,
		TotalResults: subscriptions.PageInfos.TotalResults,
		ResultPerPage: subscriptions.PageInfos.ResultsPerPage,
	}

	jsonBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	return jsonBody
}















//
//func HandleNextPage(w http.ResponseWriter, r *http.Request) {
//	var page Page
//	subscriptions, err := querySubs("https://www.googleapis.com/youtube/v3/subscriptions?access_token=%v&part=snippet&maxResults=10&pageToken=CAoQAA&mine=true")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	for _, p := range subscriptions.Items {
//		c := &payload.User{}
//
//		c, err := c.GetItemInfo(p); if err != nil {
//			fmt.Println(err.Error())
//		}
//		page.AllSubscription = append(page.AllSubscription, c)
//	}
//
//	t, err := getTemplateHTML("./html/page.html"); if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Fprintf(w, htmlHome)
//	t.Execute(w, page.AllSubscription)
//}
//


//func getTemplateHTML(filePath string) (*template.Template, error) {
//	t, err := template.ParseFiles(filePath)
//	if err != nil {
//		return &template.Template{}, err
//	}
//	return t, nil
//}