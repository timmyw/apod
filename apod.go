package apod

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func extractImageUrl(body io.ReadCloser) string {
	z := html.NewTokenizer(body)

	var lastToken html.Token

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.StartTagToken:

			t := z.Token()
			if t.Data == "img" {

				// If we have a parent that's an anchor tag, try and see if there's a href to use
				// Because...
				// "Clicking on the picture will download the highest resolution version available"
				if lastToken.Data == "a" {
					for _, atts := range lastToken.Attr {
						if atts.Key == "href" {
							return atts.Val
						}
					}
					// else use the image source
				} else {
					for _, a := range t.Attr {
						if a.Key == "src" {
							return a.Val
						}
					}
				}
				lastToken = t
			}
		}

	}

}

func getImage(url string, output string) {
	out, err := os.Create(output)
	if err != nil {
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
}

func generateOutputPath(outputDir string, imagePath string) string {
	_, file := filepath.Split(imagePath)
	return fmt.Sprintf("%s/%s", outputDir, file)
}

func ApodDownload(outputDir string, ds string) string {
	htmlUrl := fmt.Sprintf("https://apod.nasa.gov/apod/ap%s.html", ds)

	resp, err := http.Get(htmlUrl)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	defer resp.Body.Close()

	imagePath := extractImageUrl(resp.Body)
	imgUrl := fmt.Sprintf("https://apod.nasa.gov/apod/%s", imagePath)
	outputPath := generateOutputPath(outputDir, imagePath)
	getImage(imgUrl, outputPath)

	return outputPath
}

func ApodDownloadLatest(outputDir string) {
	fmt.Print(ApodDownload(outputDir, time.Now().Format("060102")))
}
