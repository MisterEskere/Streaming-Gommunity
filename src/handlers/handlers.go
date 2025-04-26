package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MisterEskere/Streaming-Gommunity/src/engines"
	"github.com/MisterEskere/Streaming-Gommunity/src/tmdb"
)

type StreamingSource struct {
	Service string
	Link    string
	Type    string
}

// TrendingMoviesHandler handles the request for trending movies of the day
// Note: This function now returns http.HandlerFunc to capture the TmdbClient
func TrendingMoviesHandler(TmdbClient *tmdb.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the trending function to get the trending movies
		trendingMovies, err := TmdbClient.TrendingMovies()
		if err != nil {
			http.Error(w, "Failed to fetch trending movies", http.StatusInternalServerError)
			return
		}

		// Set the response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode the trending movies as JSON and write to the response
		if err := json.NewEncoder(w).Encode(trendingMovies); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		// No return statement needed here for w and r
	}
}

// TrendingSeriesHandler handles the request for trending TV series of the day
// Note: This function now returns http.HandlerFunc to capture the TmdbClient
func TrendingSeriesHandler(TmdbClient *tmdb.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the trending function to get the trending TV series
		trendingSeries, err := TmdbClient.TrendingSeries()
		if err != nil {
			http.Error(w, "Failed to fetch trending series", http.StatusInternalServerError)
			return
		}

		// Set the response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode the trending series as JSON and write to the response
		if err := json.NewEncoder(w).Encode(trendingSeries); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		// No return statement needed here for w and r
	}
}

// Return the streaming URLs of the requested movie (later on implement series)
func GetStreamingLinksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the query parameters needed to find the correct element on the streaming sites
	queryParams := r.URL.Query()
	Slug := queryParams.Get("Slug")

	engines.GetStreamingCommunityLink(Slug)
}
