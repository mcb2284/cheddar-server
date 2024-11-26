package main

import (
	"github.com/cheddar/api"
	"github.com/cheddar/database"
	_ "github.com/go-sql-driver/mysql"
)

// album represents data about a record album.


type user struct {
	ID int `json:"user_id"`
	User string `json:"user_name"`
	First string `json:"first_name"`
	Last string `json:"last_name"`
}


func main(){

	database.Database()
	api.Server()

}

