package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/PuerkitoBio/goquery"
)

func errEvent(err string) []models.Event {
	log.Fatal(err)
	return []models.Event{models.Event{
		Title:   "error",
		Date:    "error",
		Artists: "error",
		Genre:   "error",
		IsImage: false,
	}}
}

func ScrapePage(url string) (*goquery.Document, []models.Event) {
	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		return nil, errEvent(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errEvent(fmt.Sprintf("Failed to fetch page: %d %s", res.StatusCode, res.Status))
	}

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, errEvent(err.Error())
	}

	return doc, nil

}
