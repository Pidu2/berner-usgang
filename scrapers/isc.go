package scrapers

import (
	"regexp"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeISC(url string) ([]models.Event, error) {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".event_preview").Each(func(i int, eventItem *goquery.Selection) {
		// ISC has lots of whitespaces, remove them
		reLeading := regexp.MustCompile(`^\s+`)
		reTrailing := regexp.MustCompile(`\s+$`)

		// Extract event date
		eventDate := eventItem.Find(".event_title_date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".event_title_title").Text()

		// Extract genre
		genre := eventItem.Find(".event_title_info_mobile").Text()

		evList = append(evList, models.Event{
			Title:   reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventTitle, ""), ""),
			Date:    reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventDate, ""), ""),
			Artists: "", // ISC has no artists
			Genre:   reTrailing.ReplaceAllString(reLeading.ReplaceAllString(genre, ""), ""),
			IsImage: false,
		})
	})
	return evList, nil
}
