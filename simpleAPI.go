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

// We'll just have a map from titles to a Quote
//var quotes map[string] Quote

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
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
	fmt.Println(port)

	// Load the quotes from the JSON file
	jsonFile, err := os.Open("ronSwansonQuotes.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ronSwansonQuotes.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var quotesArr Quotes
	quotes := make(map[int]Quote)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'quotesArr' which we defined above
	json.Unmarshal(byteValue, &quotesArr)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	// for _, quote := range quotes {
	// 	fmt.Println("Quote Title: " + quote.Title)
	// 	fmt.Println("Quote: " + quote.Quote)
	// }
	for i := 0; i < len(quotesArr.Quotes); i++ {
		fmt.Printf("Quote Id: %d\n", quotesArr.Quotes[i].ID)
		fmt.Println("Quote Title: " + quotesArr.Quotes[i].Title)
		fmt.Println("Quote: " + quotesArr.Quotes[i].Quote)
		quotes[quotesArr.Quotes[i].ID] = quotesArr.Quotes[i]
	}

	// Set up API
	r := gin.Default()
	r.LoadHTMLFiles("newQuote.html")

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/v1/quotes", func(c *gin.Context) {
		//var allQuotes = gin.H{"quotes": quotesArr}
		fmt.Println("*************")
		for _, quote := range quotes {
			fmt.Println(quote.Title)
			fmt.Println(quote.Quote)
			//allQuotes[strconv.Itoa(quote.ID)] = gin.H{"title": quote.Title, "quote": quote.Quote}
		}
		fmt.Println("*************")
		//fmt.Println(allQuotes)
		//fmt.Println(len(allQuotes))
		c.JSON(http.StatusOK, quotesArr)
	})

	r.GET("/v1/quote/new", newQuote)

	r.GET("/v1/quote", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err == nil {
			if _, ok := quotes[id]; ok {
				c.JSON(http.StatusOK, gin.H{"quote": quotes[id]})
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
		q := Quote{len(quotesArr.Quotes) + 1, title, quote}
		quotesArr.Quotes = append(quotesArr.Quotes, q)
		quotes[q.ID] = q
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
