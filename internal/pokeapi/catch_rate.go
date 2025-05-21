package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchRate(pageURL string) (PokeStats, error) {
	url := baseURL + "/pokemon/" + pageURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeStats{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeStats{}, err
	}

	pokeInfo := PokeStats{}

	err = json.Unmarshal(dat, &pokeInfo)
	if err != nil {
		return PokeStats{}, nil
	}

	return pokeInfo, nil
}
