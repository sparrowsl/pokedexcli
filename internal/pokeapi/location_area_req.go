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
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, nil
	}

	LocationAreasResp := LocationAreasResponse{}
	// unmarshall the data into Go struct
	err = json.Unmarshal(data, &LocationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, nil
	}

	return LocationAreasResp, nil
}
