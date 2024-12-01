package api

import (
	"net/http"

	"github.com/cheddar/database"
	"github.com/cheddar/types"
	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func Server(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)


	router.POST("/user", createUser)
	router.GET("/user/:id", getUser)
	router.GET("/name/:name", getUserByName)
	router.PATCH("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run("localhost:8080")
}


func createUser(c *gin.Context){
	var newUser types.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	database.CreateUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
func getUser(c *gin.Context){

	user := database.GetUser(c.Param("id"))

	c.IndentedJSON(http.StatusOK, user)
}

func getUserByName(c *gin.Context){

	user := database.GetUserByName(c.Param("name"))

	c.IndentedJSON(http.StatusOK, user)

}

func updateUser(c *gin.Context){
	return
}
func deleteUser(c *gin.Context){

	database.DeleteUser(c.Param("id"))

	c.IndentedJSON(http.StatusOK, "Deleted")
}


func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	} 

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _ , a := range albums {
        if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}