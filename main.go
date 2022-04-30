package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/cavaliergopher/grab/v3"
)

const webpage = "https://saividya.ac.in/"

func main() {
	resp, err := http.Get(webpage)

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		fmt.Println("Error creating document", err)
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgPath, _ := s.Attr("src")

		if strings.HasPrefix(imgPath, "http") {
			_, err := grab.Get(".", imgPath)

			if err != nil {
				fmt.Println("Error downloading image", err)
			} else {
				fmt.Println("Image downloaded")
			}
		} else {
			if strings.HasPrefix(imgPath, "//") {
				scheme := getScheme(webpage)
				if scheme != "" {
					_, err := grab.Get(".", scheme+":"+imgPath)
					if err != nil {
						fmt.Println("Error downloading the image")
						return
					}
				}
			}
		}
	})

}

func getScheme(link string) string {
	u, err := url.Parse(link)

	if err != nil {
		fmt.Println("Could not parse web link", err)
		return ""
	}
	return u.Scheme
}
