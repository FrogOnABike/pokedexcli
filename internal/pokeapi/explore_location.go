package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locName string) (respLocationDetail, error) {
	pageURL := baseURL + "/location-area/" + locName

	cacheData, exists := c.clientCache.Get(pageURL)
	if exists {
		fmt.Println("***Retrieving from cache***")
		locationDetail := respLocationDetail{}

		if err := json.Unmarshal(cacheData, &locationDetail); err != nil {
			return respLocationDetail{}, err
		}
		return locationDetail, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return respLocationDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return respLocationDetail{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return respLocationDetail{}, err
	}

	locationDetail := respLocationDetail{}
	c.clientCache.Add(pageURL, body)

	if err := json.Unmarshal(body, &locationDetail); err != nil {
		return respLocationDetail{}, err
	}

	return locationDetail, nil
}
