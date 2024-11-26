package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

type user struct {
	ID int `json:"user_id"`
	User string `json:"user_name"`
	First string `json:"first_name"`
	Last string `json:"last_name"`
}

var db *sql.DB

func Database(){
	var err error
	db, err = sql.Open("mysql", "root:password@/cheddar_db")

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func GetUser(id string) string {
	return getUser(db, id)
}

func getUser(db *sql.DB, id string) string {
	db, err := sql.Open("mysql", "root:password@/cheddar_db")

	if err != nil { 
		panic(err)
	}
	id_int, err := strconv.Atoi(id)

	if err != nil { 
		panic(err)
	}

	result, err := db.Query("SELECT * FROM users WHERE user_id = ?", id_int)
	
	if err != nil { 
		panic(err)
	}


	var user_id int
	var user string
	var first string
	var last string


	if result.Next(){
		err = result.Scan(&user_id, &user, &first, &last)
	} else {
		fmt.Println("Nothing found here ...")
	}


	if err != nil {
		panic(err)
	}

	fmt.Printf("Id: %d User: %s First: %s Last: %s\n", user_id, user, first, last)
	
	defer db.Close()
	return user
}