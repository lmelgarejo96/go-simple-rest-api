package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var languages = `[{"code":"english","name":"English (inglés)"},{"code":"spanish","name":"Español de España"},{"code":"french","name":"Français (francés)"},{"code":"italian","name":"Italiano (italiano)"},{"code":"german","name":"Deutsch (alemán)"},{"code":"greek","name":"Ελληνικά (griego)"},{"code":"koreana","name":"한국어 (coreano)"},{"code":"schinese","name":"简体中文 (chino simplificado)"},{"code":"tchinese","name":"繁體中文 (chino tradicional)"},{"code":"russian","name":"Русский (ruso)"},{"code":"thai","name":"ไทย (tailandés)"},{"code":"japanese","name":"日本語 (japonés)"},{"code":"portuguese","name":"Português (portugués)"},{"code":"brazilian","name":"Português-Brasil (portugués de Brasil)"},{"code":"polish","name":"Polski (polaco)"},{"code":"danish","name":"Dansk (danés)"},{"code":"dutch","name":"Nederlands (holandés)"},{"code":"finnish","name":"Suomi (finés)"},{"code":"norwegian","name":"Norsk (noruego)"},{"code":"swedish","name":"Svenska (sueco)"},{"code":"czech","name":"Čeština (checo)"},{"code":"hungarian","name":"Magyar (húngaro)"},{"code":"romanian","name":"Română (rumano)"},{"code":"bulgarian","name":"Български (búlgaro)"},{"code":"turkish","name":"Türkçe (turco)"},{"code":"ukrainian","name":"Українська (ucraniano)"},{"code":"vietnamese","name":"Tiếng Việt (vietnamita)"},{"code":"latam","name":"Español de Latinoamérica"}]`

type Lang struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func getLanguages(c *gin.Context) {
	var responseLanguages []Lang
	json.Unmarshal([]byte(languages), &responseLanguages)

	fmt.Println("data", responseLanguages)

	c.IndentedJSON(http.StatusOK, responseLanguages)
}
