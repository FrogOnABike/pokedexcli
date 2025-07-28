package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (respLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	cacheData, exists := c.clientCache.Get(url)
	if exists {
		locationsResp := respLocationArea{}

		if err := json.Unmarshal(cacheData, &locationsResp); err != nil {
			return respLocationArea{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return respLocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return respLocationArea{}, err
	}

	locationsResp := respLocationArea{}
	c.clientCache.Add(url, body)

	if err := json.Unmarshal(body, &locationsResp); err != nil {
		return respLocationArea{}, err
	}

	return locationsResp, nil
}
