package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseURL := "https://example.com" // Replace this with the URL you want to scrape
	downloadPath := "./downloads"    // Replace this with the desired download path

	err := os.MkdirAll(downloadPath, 0755)
	if err != nil {
		fmt.Println("Error creating download directory:", err)
		return
	}

	doc, err := scrapeDownloads(baseURL)
	if err != nil {
		fmt.Println("Error scraping downloads:", err)
		return
	}

	doc.Find("a[href$='pdf']").Each(func(i int, s *goquery.Selection) {
		downloadFile(baseURL, downloadPath, s)
	})

	doc.Find("a[href$='xlsx']").Each(func(i int, s *goquery.Selection) {
		downloadFile(baseURL, downloadPath, s)
	})

	// Add more selectors for other file types if needed
}

func scrapeDownloads(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func downloadFile(baseURL, downloadPath string, selection *goquery.Selection) {
	href := selection.AttrOr("href", "")

	if !strings.HasPrefix(href, "http") {
		u, err := url.Parse(baseURL)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}
		href = u.ResolveReference(&url.URL{Path: href}).String()
	}

	fmt.Println("Downloading:", href)

	resp, err := http.Get(href)
	if err != nil {
		fmt.Println("Error downloading:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Status code:", resp.Status)
		return
	}

	fileName := filepath.Base(href)
	downloadDir := downloadPath
	out, err := os.Create(filepath.Join(downloadDir, fileName))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Println("File downloaded:", fileName)
}
