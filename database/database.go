package database

import (
	"database/sql"
	"fmt"
)

func database(){

	//create db obj
	db, err := sql.Open("mysql", "root:password@/cheddar_db")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	result, err := db.Query("SELECT * FROM users")

	if err != nil { 
		panic(err)
	}

	fmt.Print(result)

	fmt.Print("Pong\n")

	defer db.Close()
}