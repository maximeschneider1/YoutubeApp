package main

import (
	"YSS/config"
	"YSS/handler"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"

)

var configPath = "/Users/max/go/src/YSS/config/DBconfig.json"
var GlobalDB *sql.DB

type server struct {
	myDB *sql.DB
}

func main() {

	db, err := config.ReturnDB(configPath); if err != nil {
		fmt.Println(err.Error())
	}
	GlobalDB = db

	getUserInfo()

	fmt.Println("Starting Web Server...")
	handler.StartWebServer()

}

type User struct {
	ID int
}

func getUserInfo() User {
	var userConnexion User
	_, _, err := isUserInDB(GlobalDB, 1); if err != nil {
		fmt.Println(err.Error())
		// STOP
	}
	return userConnexion

	//if exist == true && userID != 0 {
	//	// get user info
	//	row := GlobalDB.QueryRow("SELECT user_id FROM users_info WHERE user_id=$1;", userID)
	//	fmt.Printf("User %v existe", userID)
	//	return userConnexion
	//} else {
	//	fmt.Printf("User %v n'existe pas", userID)
	//	// post user info
	//
	//	return userConnexion
	//}
}

func isUserInDB(db *sql.DB, userID int) (bool, int, error){
	var queryResult User

	row := db.QueryRow("SELECT user_id FROM users_info WHERE user_id=$1;", userID)

	err := row.Scan(&queryResult.ID); if err != nil {
		if err.Error() == "sql: no rows in result set"{
			return false, userID, nil
		} else {
			fmt.Println(err.Error())
			return false, 0, err
		}
	}

	return true, userID, nil
}