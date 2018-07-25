package apod

import (
	"fmt"
	"time"
	"net/http"
	"io"
	"os"
	"golang.org/x/net/html"
	"path/filepath"
)

func extractImageUrl(body io.ReadCloser) string {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "img" {
				for _, a := range t.Attr {
					if a.Key == "src" {
						return a.Val
					}
				}
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

func generateOutputPath(imagePath string) string {
	_, file := filepath.Split(imagePath)
	return fmt.Sprintf("./%s", file)
}

func ApodDownload(d string) {
	htmlUrl := fmt.Sprintf("https://apod.nasa.gov/apod/ap%s.html", d)

	resp, err := http.Get(htmlUrl)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()

	imagePath := extractImageUrl(resp.Body)
	imgUrl := fmt.Sprintf("https://apod.nasa.gov/apod/%s", imagePath)
	outputPath := generateOutputPath(imagePath)
	getImage(imgUrl, outputPath)
}

func ApodDownloadLatest() {
	ApodDownload(time.Now().Format("060102"))
}
