package main

import (
	"fmt"
	"log"
	"strings"
)

func logOnErr(e error) {
	if e != nil {
		log.Println("[ERROR FOUND]:\t", e)
	}
}

func logFatalOnErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

func custom_log(f, m string, args ...any) {
	fn := strings.ToUpper(f)      // file name
	mf := fmt.Sprintf(m, args...) // message formatted
	fmt.Printf("[%s]\t%s\n", fn, mf)
	log.Printf("[%s]\t%s\n", fn, mf)
}
