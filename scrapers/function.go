package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
)

// Define a type for scraper functions
type ScraperFunc func(url string, limit int) ([]models.Event, error)
