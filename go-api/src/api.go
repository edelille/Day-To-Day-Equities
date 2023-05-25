package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Temporary
const (
	PORTNO   = ":12000"
	BASE_URL = "https://api.polygon.io/v2"
)

func handle_noargs(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Day to Day Equities API written in GO")
}

func handle_daily_quote(c *gin.Context) {
	r := HTTP_CLIENT

	var inp QuoteJSON
	c.ShouldBindJSON(&inp)

	quoteURI := fmt.Sprintf("%s/aggs/ticker/%s/range/1/minute/%s/%s?adjusted=true&apiKey=%s",
		BASE_URL,
		inp.Ticker,
		inp.Date,
		inp.Date,
		API_KEY,
	)

	fmt.Println(quoteURI)

	resp, err := r.R().
		Get(quoteURI)
	logOnErr(err)

	if !resp.IsSuccessState() {
		respStr, err := resp.ToString()
		fmt.Println("uh oh: ", resp.GetStatusCode(), respStr, err)
	} else {
		fmt.Println("Result: ", resp.String())
	}

}

func main_api() {
	fmt.Println("Starting main api routine")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "OPTIONS", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Referrer", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Hour * 12,
	}))

	// Set up API endpoints
	r.GET("/", handle_noargs)
	r.POST("/daily", handle_daily_quote)

	custom_log("api", "Starting API on port %s\n", PORTNO)
	r.Run(PORTNO)
}
