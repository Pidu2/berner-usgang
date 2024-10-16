package scrapers

import (
	"strings"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeCafete(url string, limit int) ([]models.Event, error) {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".event").Each(func(i int, eventItem *goquery.Selection) {
		if len(evList) == limit {
			return
		}
		// Extract event date
		eventDate := eventItem.Find(".date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".title").Text()

		// Extract artist
		artists := ""
		eventItem.Find(".acts").Contents().Each(func(j int, actItem *goquery.Selection) {
			if goquery.NodeName(actItem) == "br" {
				artists += ", "
			} else {
				artists += actItem.Text()
			}
		})

		// Extract genre
		genre := eventItem.Find(".style").Text()
		genre = strings.Replace(genre, "Style: ", "", 1)

		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: artists,
			Genre:   genre,
			IsImage: false,
		})
	})
	return evList, nil
}
