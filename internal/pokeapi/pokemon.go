package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		resp := RespPokemon{}
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return RespPokemon{}, err
		}
		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemon := RespPokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
