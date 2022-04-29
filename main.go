package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

//var quote [] quote = make([]quote, 0) - zero quotes
var quotes []quote = []quote{
	{"Learn to Lead", "SVIT"},
	{"1436", "no"},
	{"teyhgc", "alpha"},
	{"hdjwjhd", "HH"},
}

func main() {
	rand.Seed(time.Now().Unix())
	api := echo.New() //newinstance,create new server-echo
	//endpoint
	api.GET("/quotes", getQuotes) //handler for this particular endpoint
	api.GET("/quotes/random", getRandomQuote)
	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":" + port)
	//localhost:9090
}
func getQuotes(c echo.Context) error {
	//return c.String(200, "Hello world - Sample Program")
	return c.JSON(http.StatusOK, quotes)
}
func getRandomQuote(c echo.Context) error {

	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}
