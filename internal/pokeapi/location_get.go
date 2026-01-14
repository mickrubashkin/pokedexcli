package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationArea(areaName string) (LocationAreaDetail, error) {
	url := baseURL + "/location-area/" + areaName

	var locationAreaDetail LocationAreaDetail
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &locationAreaDetail); err != nil {
			return LocationAreaDetail{}, err
		}
		return locationAreaDetail, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaDetail{}, fmt.Errorf("bad status: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	err = json.Unmarshal(data, &locationAreaDetail)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	c.cache.Add(url, data)

	return locationAreaDetail, nil
}
