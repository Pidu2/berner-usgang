package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pidu2/berner-usgang/api"
)

func main() {
	// Register the routes
	http.HandleFunc("/scraper", api.HandleGetScrapers)
	http.HandleFunc("/scraper/", api.HandleScrape) // Note the trailing slash for dynamic scraper name

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
