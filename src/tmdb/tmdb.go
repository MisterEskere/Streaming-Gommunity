package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiReadAccessToken string
	httpClient         *http.Client
}

type TrendingElement struct {
	BackdropPath     string  `json:"backdrop_path"`
	ID               int     `json:"id"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	OriginalLanguage string  `json:"original_language"`
	GenreIDs         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type TrendingMovie struct {
	TrendingElement
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	ReleaseDate   string `json:"release_date"`
}

type TrendingSerie struct {
	TrendingElement
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	FirstAirDate string `json:"first_air_date"`
}

func NewClient(token string) *Client {
	return &Client{
		apiReadAccessToken: token,
		httpClient:         &http.Client{},
	}
}

func (c *Client) tmdbGetRequest(url string) ([]byte, error) {

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return nil, fmt.Errorf("error creating request: %v", error)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiReadAccessToken)

	res, error := c.httpClient.Do(req)
	if error != nil {
		return nil, fmt.Errorf("error making request: %v", error)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("status code: %v, response body: %s", res.StatusCode, string(body))
	}

	responseBody, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, fmt.Errorf("error reading response body: %v", error)
	}

	return responseBody, error
}

func (c *Client) TrendingMovies() ([]TrendingMovie, error) {
	url := "https://api.themoviedb.org/3/trending/movie/day?language=it-IT"

	responseBody, err := c.tmdbGetRequest(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching trending movies: %v", err)
	}

	var response struct {
		Results []TrendingMovie `json:"results"`
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing trending movies response: %v", err)
	}

	return response.Results, nil
}

func (c *Client) TrendingSeries() ([]TrendingSerie, error) {
	url := "https://api.themoviedb.org/3/trending/tv/day?language=it-IT"

	responseBody, err := c.tmdbGetRequest(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching trending series: %v", err)
	}

	var response struct {
		Results []TrendingSerie `json:"results"`
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing trending series response: %v", err)
	}

	return response.Results, nil
}
