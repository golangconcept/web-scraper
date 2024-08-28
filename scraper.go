package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	args := os.Args

	url := args[1]
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a respose from", r.Request.URL)
	})

	collector.Visit(url)
}
