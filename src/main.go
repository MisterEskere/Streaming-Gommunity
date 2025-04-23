package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// Set up the web server
	http.HandleFunc("/trendingMovies", TrendingMoviesHandler(TmdbClient))
	http.HandleFunc("/trendingSeries", TrendingSeriesHandler(TmdbClient))

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
