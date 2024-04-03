package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := baseUrl + "/location-area"

	if pageURL != nil {
		endpoint = *pageURL
	}

	// check the cache if the current url has a value
	data, ok := client.cache.Get(endpoint)
	if ok {
		// if cache exists for this URL, return the data instead of doing a network request
		LocationAreasResp := LocationAreasResponse{}
		// unmarshall the data into Go struct
		err := json.Unmarshal(data, &LocationAreasResp)
		if err != nil {
			return LocationAreasResponse{}, nil
		}

		return LocationAreasResp, nil
	}

	// Make a GET request to the endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	// Executes the request made above
	response, err := client.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, nil
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	// Read data from response body
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, nil
	}

	LocationAreasResp := LocationAreasResponse{}
	// unmarshall the data into Go struct
	err = json.Unmarshal(data, &LocationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, nil
	}

	// before returning the data,
	// save current data inside the cache
	client.cache.Add(endpoint, data)

	return LocationAreasResp, nil
}

func (client *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := baseUrl + "/location-area/" + locationAreaName

	// check the cache if the current url has a value
	data, ok := client.cache.Get(endpoint)
	if ok {
		// if cache exists for this URL, return the data instead of doing a network request
		locationArea := LocationArea{}
		// unmarshall the data into Go struct
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, nil
		}

		return LocationArea{}, nil
	}

	// Make a GET request to the endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// Executes the request made above
	response, err := client.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, nil
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	// Read data from response body
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, nil
	}

	locationArea := LocationArea{}
	// unmarshall the data into Go struct
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, nil
	}

	// before returning the data,
	// save current data inside the cache
	client.cache.Add(endpoint, data)

	return locationArea, nil
}
