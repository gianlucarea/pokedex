package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		resp := RespLocationArea{}
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationArea := RespLocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)
	return locationArea, nil
}
