package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (LocationAreaListResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationsResp := LocationAreaListResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			fmt.Println(err)
			return LocationAreaListResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaListResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaListResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaListResponse{}, fmt.Errorf("bad status: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaListResponse{}, err
	}

	locationsResp := LocationAreaListResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		fmt.Println(err)
		return LocationAreaListResponse{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
