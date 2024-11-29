package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumURL = "http://localhost:4444/wd/hub"
)

func main() {
	caps := selenium.Capabilities{"browserName": "chrome"}

	driver, err := selenium.NewDriver(selenium.Remote("webdriver.chrome.driver", caps))
	if err != nil {
		log.Fatal("Error creating driver:", err)
	}
	defer driver.Quit()

	err = driver.Get("https://example.com")
	if err != nil {
		log.Fatal("Error getting URL:", err)
	}

	downloadLinks, err := getDownloadLinks(driver)
	if err != nil {
		log.Fatal("Error getting download links:", err)
	}

	for _, link := range downloadLinks {
		downloadFile(driver, link)
	}
}

func getDownloadLinks(driver selenium.WebDriver) ([]string, error) {
	var links []string

	elements, err := driver.FindElements(selenium.ByCSSSelector("a[href$=.pdf]").Add(selenium.ByCSSSelector("a[href$=.xlsx]")))
	if err != nil {
		return nil, err
	}

	for _, element := range elements {
		link, err := element.Attribute("href")
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func downloadFile(driver selenium.WebDriver, url string) {
	fmt.Println("Downloading:", url)

	_, err := driver.Get(url)
	if err != nil {
		log.Fatal("Error downloading:", err)
		return
	}

	// Wait for the download to complete
	// You may need to adjust this wait time based on the file size
	time.Sleep(5 * time.Second)

	downloadPath, err := getDownloadPath(driver)
	if err != nil {
		log.Fatal("Error getting download path:", err)
		return
	}

	filename := strings.Split(url, "/")
	src := filepath.Join(downloadPath, filename[len(filename)-1])
	dst := filepath.Join("downloads", filename)

	err = os.Rename(src, dst)
	if err != nil {
		log.Fatal("Error moving downloaded file:", err)
		return
	}

	fmt.Println("Downloaded:", filename)
}

func getDownloadPath(driver selenium.WebDriver) (string, error) {
	prefs, err := driver.Manage().GetPreferences()
	if err != nil {
		return "", err
	}

	downloadPath, ok := prefs["download.default_directory"]
	if !ok {
		return "", fmt.Errorf("download.default_directory preference not found")
	}

	return downloadPath.(string), nil
}
