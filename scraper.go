package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

func main() {

	var products []Product
	// args := os.Args

	// _url := args[1]
	collector := colly.NewCollector(
	// colly.AllowedDomains("www.scrapingcourse.com"),
	)

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// triggered when the scraper encounters an error
	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a respose from", r.Request.URL)
	})

	// triggered when a CSS selector matches an element
	// collector.OnHTML("a", func(e *colly.HTMLElement) {
	// 	// printing all URLs associated with the <a> tag on the page
	// 	fmt.Println("%v", e.Attr("href"))
	// })

	collector.OnHTML("li.product", func(e *colly.HTMLElement) {

		// initialize a new Product instance
		product := Product{}

		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")

		// add the product instance with scraped data to the list of products
		products = append(products, product)

	})
	// triggered once scraping is done (e.g., write the data to a CSV file)
	collector.OnScraped(func(r *colly.Response) {
		// open the CSV file
		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initialize a file writer
		writer := csv.NewWriter(file)

		// write the CSV headers
		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

		// write each product as a CSV row
		for _, product := range products {
			// convert a Product to an array of strings
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}

			// add a CSV record to the output file
			writer.Write(record)
		}
		defer writer.Flush()
		fmt.Println(r.Request.URL, " scraped!")
	})
	collector.Visit("https://www.scrapingcourse.com/ecommerce")
}
