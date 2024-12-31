package main

import (
	"github.com/cheddar/api"
	"github.com/cheddar/database"
	_ "github.com/go-sql-driver/mysql"
)



func main(){
	
	db := database.Database()
	api.Server()

	defer db.Close()
}

