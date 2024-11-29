package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://example.com" // Replace with the URL containing download links

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

	downloadLinks := doc.Find("a[href$=.pdf]")

	for i, _ := range downloadLinks.Nodes {
		link := downloadLinks.Eq(i).AttrOr("href", "")
		if link == "" {
			continue
		}
		downloadFile(link)
	}
}

func downloadFile(url string) {
	fmt.Println("Downloading:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Status code:", resp.Status)
		return
	}

	filename := strings.Split(url, "/")
	out, err := os.Create(filepath.Join("downloads", filename[len(filename)-1]))
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

	fmt.Println("Downloaded:", filename)
}
