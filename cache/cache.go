package cache

import (
	"time"

	"github.com/Pidu2/berner-usgang/globals"
	"github.com/Pidu2/berner-usgang/models"
	"github.com/patrickmn/go-cache"
)

var (
	CacheInstance = cache.New(4*time.Hour, 10*time.Minute)
)

func ScrapeWithCache(scraper globals.Scraper, limit int) ([]models.Event, error) {
	// Check if we have cached data for this key
	cacheKey := scraper.URL
	if cachedData, found := CacheInstance.Get(cacheKey); found {
		cachedItems := cachedData.([]models.Event)

		// If the cache contains enough items, return the number of requested items
		if len(cachedItems) >= limit {
			return cachedItems[:limit], nil
		}
	}

	// If not enough cached items, scrape new data
	scrapedItems, err := scraper.Function(scraper.URL, limit)
	if err != nil {
		return nil, err
	}

	// Cache the new data (overwrite existing cache for this key)
	CacheInstance.Set(cacheKey, scrapedItems, cache.DefaultExpiration)

	return scrapedItems, nil
}
