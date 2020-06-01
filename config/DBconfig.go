package config

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


type DBconfig struct {

	URL string `json:"url"`

	//Name     string `json:"dbname"`
	//Host     string `json:"host"`
	//User     string `json:"user"`
	//Database string `json:"database"`
	//Port     int `json:"port"`
	//Password string `json:"password"`
}


// ReturnDB reads json config file and returns an DB connection
func ReturnDB(configPath string) (*sql.DB, error) {

	dbc, err := readConfig(configPath); if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", dbc.URL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to DB")

	return db, nil
}

// readConfig reads json client secret to return db config
func readConfig(configPath string ) (DBconfig, error) {

	jsonFile, err := os.Open(configPath); if err != nil {
		log.Printf("Error openning file for config path : %v, error : %v", configPath, err.Error())
		return DBconfig{}, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var DBconfig = DBconfig{}
	json.Unmarshal(byteValue, &DBconfig)

	return DBconfig, nil
}