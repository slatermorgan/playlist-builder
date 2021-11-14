package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type config struct {
	baseURL string
}

type Client struct {
	config config
}

func (c *Client) Get(requestModel interface{}, url string) (interface{}, error) {
	resp, err := http.Get(c.config.baseURL + url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(bodyBytes), requestModel)
	if err != nil {
		return nil, err
	}
}
