package scrapers

import (
	"regexp"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeStellwerk(url string) ([]models.Event, error) {
	reLeading := regexp.MustCompile(`^\s+`)
	reTrailing := regexp.MustCompile(`\s+$`)
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".post").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventItem.Find(".post__date").First().Text(), ""), "")

		// Extract event title
		eventTitle := eventItem.Find("h2").Text()

		// Extract artist
		artists := eventItem.Find("h3").Text()

		// Extract genre
		genre := ""
		eventItem.Find(".post__tag").Each((func(j int, genreItem *goquery.Selection) {
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

	return evList, nil
}
