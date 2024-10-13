package scrapers

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeKapitel(url string, limit int) ([]models.Event, error) {
	evList := []models.Event{}

	currentTime := time.Now()
	year := currentTime.Year() % 100
	month := int(currentTime.Month())
	re := regexp.MustCompile(`[^a-zA-Z0-9 _\-/\.]+`)
	reLeading := regexp.MustCompile(`^\s+`)
	reTrailing := regexp.MustCompile(`\s+$`)
	spaceRe := regexp.MustCompile(` {2,}`)
	for m := range []int{0, 1} {
		loopMonth := month + m
		loopYear := year
		if loopMonth == 13 {
			loopYear += 1
			loopMonth = 1
		}
		loopUrl := fmt.Sprintf("%s%d-%02d", url, loopYear, loopMonth)
		doc, err := utils.ScrapePage(loopUrl)

		if err != nil {
			return nil, err
		}

		// Find the event list and iterate over each event item
		doc.Find(".event-link").Each(func(i int, eventItem *goquery.Selection) {
			if len(evList) == limit {
				return
			}
			// Extract event date
			eventDate := reTrailing.ReplaceAllString(reLeading.ReplaceAllString(re.ReplaceAllString(eventItem.Find(".event-date-inner").Text(), ""), ""), "")

			// Extract event title
			eventTitle := reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventItem.Find(".size-medium").Text(), ""), "")

			// Extract artist
			artists := spaceRe.ReplaceAllString(reTrailing.ReplaceAllString(reLeading.ReplaceAllString(re.ReplaceAllString(eventItem.Find(".event-artist-list").Text(), ""), ""), ""), ", ")

			// Extract genre
			genre := ""
			eventItem.Find(".event-tag").Each(func(j int, genreItem *goquery.Selection) {
				genre += ", " + genreItem.Text()
			})
			if genre != "" {
				genre = genre[1:]
			}

			if eventDate == "Club" {
				return
			}
			evList = append(evList, models.Event{
				Title:   eventTitle,
				Date:    eventDate,
				Artists: artists,
				Genre:   genre,
				IsImage: false,
			})
		})
	}

	return evList, nil
}
