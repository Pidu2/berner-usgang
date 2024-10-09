package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
)

func ScrapeDeadEnd(url string) ([]models.Event, error) {
	evList := []models.Event{}
	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	programImg, _ := doc.Find("#content").First().Find("img").First().Attr("src")
	evList = append(evList, models.Event{
		Title:   programImg,
		IsImage: true,
	})

	return evList, nil
}
