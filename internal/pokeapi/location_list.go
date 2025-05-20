package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zmdelk/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, ok := cache.Get(url)

	if ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	cache.Add(url, dat)

	locationsResp := RespShallowLocations{}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil

}
