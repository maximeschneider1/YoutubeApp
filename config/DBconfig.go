package config

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"io/ioutil"
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

	dbc := readConfig(configPath)

	//dbc := "postgres://pxpgkcmmenoikj:dfb3975b63ac20366da320feee2c5859c32de6ad05c854ce90fd6f5e6c3e2cd0@ec2-54-247-181-239.eu-west-1.compute.amazonaws.com:5432/d1c5b8qlqavuau"

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=verify-full",
	//	dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Name)

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
func readConfig(configPath string ) DBconfig {

	jsonFile, err := os.Open(configPath); if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var DBconfig = DBconfig{}
	json.Unmarshal(byteValue, &DBconfig)

	return DBconfig
}