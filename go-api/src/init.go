package main

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
)

func init_logging() {
	file, err := os.OpenFile("logs/temp.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logFatalOnErr(err)

	log.SetOutput(file)
	log.Println("Log Initialized on ", time.Now().String())
}

func init_metadata() {
	godotenv.Load("local.env")

	API_KEY = os.Getenv("API_KEY")
}

func init_http_client() {
	httpClient := req.C()

	httpClient = httpClient.EnableForceHTTP1()

	httpClient = httpClient.SetTLSClientConfig(&tls.Config{
		Renegotiation: tls.RenegotiateOnceAsClient,
	})

	HTTP_CLIENT = httpClient
}

func init() {
	init_metadata()
	custom_log("init", "Successfully finished init metadata")

	init_http_client()
	custom_log("init", "Successfully finished init http client")
}
