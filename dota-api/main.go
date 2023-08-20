package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ResponseApiHeroes struct {
	Result struct {
		Data struct {
			Heroes []SimpleHeroe `json:"heroes"`
		}
	}
}

type ResponseApiHeroe struct {
	Result struct {
		Data struct {
			Heroes []interface{}
		}
	}
}

type SimpleHeroe struct {
	ID          int    `json:"id"`
	Name        string `json:"name_english_loc"`
	SkuName     string `json:"name"`
	Complexity  int    `json:"complexity"`
	NameLoc     string `json:"name_loc"`
	PrimaryAttr int    `json:"primary_attr"`
}

func fetchGetAPI(url string) ([]byte, int, error) {
	var status = 500
	defaultBytes := []byte{}
	response, err := http.Get(url)

	status = response.StatusCode

	if err != nil {
		return defaultBytes, status, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return defaultBytes, 500, err
	}

	return responseData, status, err
}

func getAllHeroes(c *gin.Context) {
	var lang = c.Query("lang")

	if len(lang) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "El parámetro 'lang' es requerido"})
		return
	}

	var apiURL = os.Getenv("API_DOTA_URL") + "/herolist?language=" + lang
	response, status, err := fetchGetAPI(apiURL)

	if err != nil {
		c.IndentedJSON(status, gin.H{"message": err.Error()})
		return
	}

	var responseHeroes ResponseApiHeroes
	json.Unmarshal(response, &responseHeroes)

	fmt.Println("data", responseHeroes)

	c.IndentedJSON(http.StatusOK, responseHeroes.Result.Data.Heroes)
}

func getHeroeById(c *gin.Context) {
	var heroId = c.Param("id")
	var lang = c.Query("lang")

	if len(lang) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "El parámetro 'lang' es requerido"})
		return
	}

	var apiURL = os.Getenv("API_DOTA_URL") + "/herodata?language=" + lang + "&hero_id=" + heroId

	response, status, err := fetchGetAPI(apiURL)

	if err != nil {
		c.IndentedJSON(status, gin.H{"message": err.Error()})
		return
	}

	var responseHeroe ResponseApiHeroe
	json.Unmarshal(response, &responseHeroe)

	if len(responseHeroe.Result.Data.Heroes) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se encontró el heroe buscado"})
		return
	}

	c.IndentedJSON(http.StatusOK, responseHeroe.Result.Data.Heroes[0])
}

func main() {
	// Permite leer las variables de entorno desde el .env
	err := godotenv.Load()

	if err != nil {
		fmt.Println("No se pudo cargar el archivo .env")
	}

	// Una forma de declarar el router (Shor variable declaration :=)
	router := gin.Default() // Inicia el router

	// Definir las rutas de la app y asignarle sus handlers
	router.GET("/languages", getLanguages)
	router.GET("/heroes", getAllHeroes)
	router.GET("/heroes/:id", getHeroeById)

	// Escuchar el app en un puerto
	router.Run("localhost:8091")
}
