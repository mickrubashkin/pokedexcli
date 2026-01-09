package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (LocationAreaListResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	return locationsResp, nil
}
