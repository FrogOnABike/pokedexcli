package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func commandMap(c *Config) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var la_JSON []respLocationArea
	if err := json.Unmarshal(body, &la_JSON); err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
