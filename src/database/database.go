package database

import (
	"database/sql"
	"fmt"
)

func Database(){

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

	fmt.Println("DB results")
	for result.Next() {
		
		var id int
		var user string
		var first string
		var last string


		err := result.Scan(&id, &user, &first, &last)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d User: %s First: %s Last: %s\n", id, user, first, last)

	}

	fmt.Print("Pong\n")

	defer db.Close()
}