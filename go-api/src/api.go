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
	var inp QuoteJSON
	QuoteURI := fmt.Sprintf("%s/aggs/ticker/%s/range/1/minute/%s/%s?adjusted=true",
		BASE_URL,
		inp.Ticker,
		inp.Date,
	)

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

	fmt.Printf("Starting API on port %s", PORTNO)
	r.Run(PORTNO)
}
