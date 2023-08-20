package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album estructura
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	// IndentedJSON permite hacer un response en formato JSON
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	var albumId = c.Param("id")

	for _, album := range albums { // Una forma de iterar, similar a un for of
		if album.ID == albumId {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	// BindJSON parsea todo el cuerpo de la Data a la struc que se utiliza
	if err := c.BindJSON(&newAlbum); err != nil {
		// AbortWithError permite generar un internal error y responder vía http, donde se le pasa un err como param
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	albums = append(albums, newAlbum) // append permite añadir items a una lista
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbum(c *gin.Context) {
	var albumId = c.Param("id")
	var idxToRemove int = -1

	for i := 0; i < len(albums); i++ {
		if albums[i].ID == albumId {
			idxToRemove = i
			break
		}
	}

	if idxToRemove == -1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se encontró el Album de ID: " + albumId})
		return
	}

	albums = append(albums[:idxToRemove], albums[idxToRemove+1:]...)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Se eliminó el Album de ID: " + albumId})
}

func main() {
	// Una forma de declarar el router (Shor variable declaration :=)
	router := gin.Default() // Inicia el router

	// Definir las rutas de la app y asignarle sus handlers
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	// Escuchar el app en un puerto
	router.Run("localhost:8090")
}
