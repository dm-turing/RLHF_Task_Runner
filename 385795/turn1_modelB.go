package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://example.com" // Replace this with the URL you want to scrape
	doc, err := scrape(url)
	if err != nil {
		log.Fatalf("Error scraping: %v", err)
	}

	scrapeData(doc)
}

func scrape(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func scrapeData(doc *goquery.Document) {
	// Example data scraping logic
	doc.Find(".product-name").Each(func(i int, sel *goquery.Selection) {
		name := sel.Text()
		fmt.Println("Product Name:", name)
	})

	doc.Find(".product-price").Each(func(i int, sel *goquery.Selection) {
		price := sel.Text()
		fmt.Println("Product Price:", price)
	})
}
