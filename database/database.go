package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/cheddar/types"
)

var db *sql.DB

func Database() *sql.DB{
	var err error
	db, err = sql.Open("mysql", "root:password@/cheddar_db")

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}
    return db
}

func GetUser(id string) string {
	return getUser(db, id)
}

func getUser(db *sql.DB, id string) string {
	
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
	
	return user
}

func CreateUser (user types.User){

	createUser(db, user)
}


func createUser(db *sql.DB, user types.User){

		stmt, err := db.Prepare("INSERT INTO users VALUES(?, ?, ?, ?)")
		randomNumber := rand.Intn(1000000)

		if err != nil {
			panic(err)
		}

		res, err := stmt.Exec(randomNumber, user.User, user.First, user.Last)
		
		if err != nil {
			panic(err)
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		fmt.Printf("ID = %d, Row = %d\n", lastId, rowCnt)
	}