package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pidu2/berner-usgang/cache"
	"github.com/Pidu2/berner-usgang/globals"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")   // Allow localhost:8080
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Specify allowed methods
}

func HandleGetScrapers(w http.ResponseWriter, r *http.Request) {
	// Return as JSON
	enableCORS(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globals.AvailableScrapers)
}

// Handle the /scraper/<scraper> endpoint - calls the scraper function
func HandleScrape(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	// Extract the scraper name from the URL path
	scraperName := r.URL.Path[len("/scraper/"):] // e.g., "dachstock", "gaskessel", "isc"

	// Check if the scraper exists
	scraper, exists := globals.AvailableScrapers[scraperName]
	if !exists || scraper.Enabled != "true" {
		http.Error(w, "Scraper not found", http.StatusNotFound)
		return
	}

	// Get the 'limit' parameter from the query string (default to 10 if not provided)
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 5 // Default limit
	}

	// Call the actual scraper function
	items, err := cache.ScrapeWithCache(scraper, limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping %s: %v", scraperName, err), http.StatusInternalServerError)
		return
	}

	// Return the scraped items as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
