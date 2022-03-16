package ngrok

import (
	"encoding/json"
	"net/http"
)

const webAddr = "http://localhost:4040"

type Client struct{}

func (c Client) Tunnels() ([]Tunnel, error) {
	resp, err := http.DefaultClient.Get(webAddr + "/api/tunnels")
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(resp.Body)
	var tr TunnelsResponse
	if err = dec.Decode(&tr); err != nil {
		return nil, err
	}
	return tr.Tunnels, nil
}

func (c Client) Requests() ([]Request, error) {
	resp, err := http.DefaultClient.Get(webAddr + "/api/tunnels")
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(resp.Body)
	var rr RequestsResponse
	if err = dec.Decode(&rr); err != nil {
		return nil, err
	}
	return rr.Requests, nil
}
