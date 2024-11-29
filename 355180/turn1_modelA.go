package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://example.com" // Replace this with the URL you want to scrape

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Status code:", resp.Status)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}
	fmt.Println(doc)
	// Now you have the parsed HTML document and can start scraping
}
