package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeRoessli(url string) []models.Event {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".page-rossli-events").Find(".event").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".event-date").Text()

		// Extract event title
		eventTitle := eventItem.Find("h2").Text()

		// Extract artist
		artists := ""

		// Extract genre
		genre := ""
		eventItem.Find("li").Each((func(j int, genreItem *goquery.Selection) {
			genre += ", " + genreItem.Text()
		}))
		// remove first comma
		if genre != "" {
			genre = genre[1:]
		}

		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: artists,
			Genre:   genre,
			IsImage: false,
		})
	})

	return evList
}
