package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zmdelk/pokedexcli/internal/pokecache"
)

func (c *Client) ListPokemon(location string, cache *pokecache.Cache) (RespFullInfo, error) {
	url := baseURL + "/location-area/" + location

	dat, ok := cache.Get(url)

	if ok {
		fullInfo := RespFullInfo{}
		err := json.Unmarshal(dat, &fullInfo)
		if err != nil {
			return RespFullInfo{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespFullInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespFullInfo{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespFullInfo{}, err
	}
	cache.Add(url, dat)

	fullInfo := RespFullInfo{}

	err = json.Unmarshal(dat, &fullInfo)
	if err != nil {
		return RespFullInfo{}, err
	}
	return fullInfo, nil
}
