package cmdlinetoolsutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gocolly/colly"
	"github.com/pkcarreno/avdm/internal/dir"
	"github.com/schollz/progressbar/v3"
)

func GetLatestVersion() (string, error) {
	// Define version number variable
	var VersionNumber string

	// Create a colly instance
	c := colly.NewCollector(
		colly.AllowedDomains("developer.android.com"),
	)

	// Create a spinner
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner

	// Define the operative system
	OS := strings.ToLower(runtime.GOOS)

	// On every a element which has href attribute call callback
	c.OnHTML("a[id=\"agree-button__sdk_"+OS+"_download\"]", func(e *colly.HTMLElement) {
		// Pick data
		linkWithVersionNumber := e.Attr("href")

		// Use regex to pick version number
		regex := regexp.MustCompile("[0-9]+")
		PickedVersionNumber := regex.FindString(linkWithVersionNumber)

		spinner.FinalMSG = PickedVersionNumber + " is the latest version.\n"
		VersionNumber = PickedVersionNumber
		spinner.Stop()
	})

	urlToVisit := "https://developer.android.com/studio"

	spinner.Suffix = " Querying " + urlToVisit + " for latest cmdline-tools version..."

	// Start the spinner
	spinner.Start()
	// Start scraping
	c.Visit(urlToVisit)

	c.Wait()
	return VersionNumber, nil
}

func DownloadPackage(version string, basepath string) (string, error) {
	// Define the operative system
	OS := strings.ToLower(runtime.GOOS)
	// Set download url
	url := "https://dl.google.com/android/repository/commandlinetools-" + OS + "-" + version + "_latest.zip"

	fileName := "commandlinetools-" + version + ".zip"
	tempPath := basepath + "temp/"
	tempFilePath := tempPath + fileName

	dir.CheckPath(tempPath)

	httpClient := http.Client{Timeout: 1 * time.Minute}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		// TODO: los errores como timeout no se estan ejecutando correctamente
		if os.IsTimeout(err) {
			log.Printf("timeout error: %v\n", err)
		}

		return "", err
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile(tempFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	fmt.Println("Downloading version " + version + " of cmdline-tools from " + url)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

	return tempFilePath, nil
}
