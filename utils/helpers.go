package utils

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func handleError(err error) error {
	log.Fatal(err)
	return err
}

func ScrapePage(url string) (*goquery.Document, error) {
	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		return nil, handleError(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, handleError(err)
	}

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, handleError(err)
	}

	return doc, nil

}
