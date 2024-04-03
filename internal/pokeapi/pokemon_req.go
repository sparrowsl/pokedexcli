package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := baseUrl + "/pokemon/" + pokemonName

	// check the cache if the current url has a value
	data, ok := client.cache.Get(endpoint)
	if ok {
		// if cache exists for this URL, return the data instead of doing a network request
		pokemon := Pokemon{}
		// unmarshall the data into Go struct
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, nil
		}

		return Pokemon{}, nil
	}

	// Make a GET request to the endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Executes the request made above
	response, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	// Read data from response body
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	// unmarshall the data into Go struct
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}

	// before returning the data,
	// save current data inside the cache
	client.cache.Add(endpoint, data)

	return pokemon, nil
}
