package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MisterEskere/Streaming-Gommunity/src/engines"
	"github.com/MisterEskere/Streaming-Gommunity/src/handlers"
	"github.com/MisterEskere/Streaming-Gommunity/src/tmdb"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the client that will be used by the handlers
	API_READ_ACCESS_TOKEN := os.Getenv("API_READ_ACCESS_TOKEN")
	TmdbClient := tmdb.NewClient(API_READ_ACCESS_TOKEN)

	// Load the engines urls (TODO, make it so that if an engine is not needed it wont be loaded)
	engines.LoadStreamingCommunityUrl()

	// Set up the web server
	http.HandleFunc("/trendingMovies", handlers.TrendingMoviesHandler(TmdbClient))
	http.HandleFunc("/trendingSeries", handlers.TrendingSeriesHandler(TmdbClient))
	http.HandleFunc("/getStreamingUrls", handlers.GetStreamingLinksHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
