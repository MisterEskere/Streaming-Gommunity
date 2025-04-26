package engines

import (
	"encoding/json"
	"os"
)

// GetSiteValue opens the sites_map.json file and returns the desired value for a given key.
func GetSiteUrl(site string) (url string) {

	file, err := os.Open("engines/sites_map.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var siteMap map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&siteMap); err != nil {
		panic(err)
	}

	url, ok := siteMap[site]
	if !ok {
		panic("site not found in sites_map.json")
	}

	return url
}
