/*
 *  A basic Ron Swanson quotes API - Andrew Emrazian
 *
 *	Usage:
 *	Start API
 *	$ go get github.com/gin-gonic/gin  # download/install API package
 *	$ go run simpleAPI.go 3005  # Optional port argument. Defaults to 3000
 *
 *	GET /v1/quotes  # get all quotes
 *	$ curl http://127.0.0.1:3005/v1/quotes
 *
 *	GET /v1/quotes/new  # new quote html form
 *	$ curl http://127.0.0.1:3005/v1/quote/new
 *
 *	GET /v1/quote?id=1  # find quote by id
 *	$ curl http://127.0.0.1:3005/v1/quote?id=1
 *
 *	POST /v1/quote  # create quote
 *	$ curl -d "title=Not%20a%20rabbit&quote=Server:Ron,%20would%20you%20like%20some%20salad?%0ARon:%20Since%20I%20am%20not%20a%20rabbit,%20no,%20I%20do%20not." \
 *	       -X POST http://127.0.0.1:3005/v1/quote
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SwansonQuotes struct {
	Quotes []SwansonQuote `json:"swansonQuotes"`
}

type SwansonQuote struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Quote string `json:"quote"`
}

func newQuote(c *gin.Context) {
	c.HTML(http.StatusOK, "newQuote.html", gin.H{"title": "New Quote"})
}

func main() {
	// Get the port from the command-line argument
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	// fmt.Println(port)

	// Load the quotes from the JSON file
	jsonFile, err := os.Open("ronSwansonQuotes.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ronSwansonQuotes.json")
	defer jsonFile.Close()

	// Load the json objects into an array and a map
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var quotesArr SwansonQuotes
	quotesMap := make(map[int]SwansonQuote)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'quotesArr' which we defined above
	json.Unmarshal(byteValue, &quotesArr)

	for i := 0; i < len(quotesArr.Quotes); i++ {
		quotesMap[quotesArr.Quotes[i].ID] = quotesArr.Quotes[i]
	}

	// Set up API
	r := gin.Default()
	r.LoadHTMLFiles("newQuote.html")

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/v1/quotes", func(c *gin.Context) {
		c.JSON(http.StatusOK, quotesArr)
	})

	r.GET("/v1/quote/new", newQuote)

	r.GET("/v1/quote", func(c *gin.Context) {
		fmt.Printf("Param id: %s\n", c.Query("id"))
		id, err := strconv.Atoi(c.Query("id"))
		if err == nil {
			if _, ok := quotesMap[id]; ok {
				c.JSON(http.StatusOK, gin.H{"quote": quotesMap[id]})
			} else {
				c.JSON(http.StatusNotFound, gin.H{"404": "not found"})
			}
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "parameter id is required"})
		}
	})

	r.POST("/v1/quote", func(c *gin.Context) {
		title := c.PostForm("title")
		quote := c.PostForm("quote")
		q := SwansonQuote{len(quotesArr.Quotes) + 1, title, quote}
		quotesArr.Quotes = append(quotesArr.Quotes, q)
		quotesMap[q.ID] = q
		c.JSON(http.StatusCreated, gin.H{
			"quote": gin.H{
				"id":    q.ID,
				"title": q.Title,
				"quote": q.Quote,
			},
		})
	})

	r.Run(":" + port)
}
